/* eslint-disable @typescript-eslint/no-explicit-any */
'use client'
import { useEffect, useState } from 'react'
import { createPublicClient, createWalletClient, custom, http, parseEventLogs, WalletClient } from 'viem'
import { baseSepolia } from 'viem/chains'
import foundnoneVrfAbi from './abi/foundnone-vrf.json'
import { wordlists } from 'bip39'

export default function Home() {
  const CONTRACT_ADDRESS: `0x${string}` = process.env.NEXT_PUBLIC_VRF_CONTRACT_ADDRESS as `0x${string}` || '0x285215e7069b10317422cd4bc8E9Cce665939545'
  const [client, setClient] = useState<WalletClient | null>(null)
  const [account, setAccount] = useState<`0x${string}`>()
  const [rand, setRand] = useState<string | null>(null)
  const [words, setWords] = useState<string[]>([])
  const [contractBalance, setContractBalance] = useState<string | null>(null)
  const [fulfillerBalance, setFulfillerBalance] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  const publicClient = createPublicClient({
    chain: baseSepolia,
    transport: http(),
  })

  useEffect(() => {
    if (client) {
      updateBalances()
    }
  }, [client, account, rand])

  async function updateBalances() {
    const contractFullBalance = await publicClient.getBalance({
      address: CONTRACT_ADDRESS,
    });
    const contractFeeBalance = await publicClient.readContract({
      address: CONTRACT_ADDRESS,
      abi: foundnoneVrfAbi,
      functionName: 'contractFeeBalance',
    }) as any;

    const aggregateFulfillerBalance = BigInt(contractFullBalance.toString()) - BigInt(contractFeeBalance.toString())
    setContractBalance(contractFeeBalance.toString())
    setFulfillerBalance(aggregateFulfillerBalance.toString())

  }

  async function connectWallet() {
    const provider = (window as any).ethereum
    if (!provider) return

    if (account) {
      await provider.request({
        method: 'wallet_requestPermissions',
        params: [{ eth_accounts: {} }],
      })
    }

    await provider.request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: `0x${baseSepolia.id.toString(16)}` }],
    })

    const accounts: string[] = await provider.request({ method: 'eth_requestAccounts' })
    const acc = accounts[0] as `0x${string}`
    setAccount(acc)

    const walletClient = createWalletClient({
      transport: custom((window as any).ethereum),
      chain: baseSepolia,
      account: acc,
    })
    setClient(walletClient)
  }

  async function requestRng() {
    if (!client || !account) return
    setLoading(true)
    try {
      const vrfFee = await publicClient.readContract({
        address: CONTRACT_ADDRESS,
        abi: foundnoneVrfAbi,
        functionName: 'requestFee',
      })
      const hash = await client.writeContract({
        account,
        address: CONTRACT_ADDRESS,
        chain: baseSepolia,
        abi: foundnoneVrfAbi,
        functionName: 'requestRng',
        value: vrfFee as any,
      })
      const receipt = await publicClient.waitForTransactionReceipt({
        hash,
        confirmations: 1,
      })
      const parsedLogs = parseEventLogs({
        logs: receipt.logs,
        abi: foundnoneVrfAbi,
        eventName: 'RngRequested',
      })[0] as any
      const requestId = parsedLogs.args.requestId
      console.log('requestId', requestId)
      await pullForEntropy(requestId)
    } catch (error) {
      console.error('Error requesting RNG:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (!rand) return
    const WORD_COUNT = 8
    const hex = BigInt(rand).toString(16)
    const bits = hex.padStart(WORD_COUNT * 4, '0')
    const chunkSize = bits.length / WORD_COUNT
    const w = Array.from({ length: WORD_COUNT }, (_, i) => {
      const slice = bits.slice(i * chunkSize, (i + 1) * chunkSize)
      const idx = parseInt(slice, 16) % wordlists.english.length
      return wordlists.english[idx]
    })
    setWords(w)
  }, [rand])

  async function pullForEntropy(requestId: string) {
    Promise.resolve(() => new Promise(resolve => setTimeout(resolve, 3000)))
    let tries = 0;
    while (tries < 10) {
      const rng = await publicClient.readContract({
        address: CONTRACT_ADDRESS,
        abi: foundnoneVrfAbi,
        functionName: 'entropies',
        args: [requestId]
      })
      if (rng) {
        setRand(rng.toString())
        break
      }
      tries++
      await new Promise(resolve => setTimeout(resolve, 1000))
    }
  }

  return (
    <div className="bg-black text-green-400 font-mono p-6 rounded-lg shadow-inner w-full h-screen">
      <div className="mb-4">
        {account ? (
          <div className='cursor-pointer' onClick={connectWallet}>&gt; Connected account: {account}</div>
        ) : (
          <div className='cursor-pointer' onClick={connectWallet}>&gt; Connect Wallet</div>
        )}
      </div>

      {contractBalance && fulfillerBalance && (
        <div className="mb-4">
          <div>&gt; Aggregate Fulfiller Balance: {(parseInt(fulfillerBalance) / 1e18).toFixed(6)} ETH</div>
          <div>&gt; Contract Fee Balance: {(parseInt(contractBalance) / 1e18).toFixed(6)} ETH</div>
        </div>
      )}

      <div className="mb-4">
        {account && (
          <div className="mt-1">&gt; <button className='cursor-pointer' onClick={requestRng}>Request RNG</button></div>
        )}
      </div>


      {rand && !loading && (
        <div className="mb-4">
          <div className='wrap-anywhere'>&gt; Entropy: {rand}</div>
          <div>&gt; Mod 100,000 / 100,000: {(parseInt(rand) % 100000) / 100000}</div>
          <div>&gt; Random words: {words.join(' ')}</div>
        </div>
      )}
      {loading && (
        <div className="mb-4">
          <div>&gt; Loading...</div>
        </div>
      )}
    </div>
  )
}
