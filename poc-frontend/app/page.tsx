/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
'use client'
import { useEffect, useRef, useState } from 'react'
import { createPublicClient, createWalletClient, custom, fallback, http, parseEventLogs, WalletClient, webSocket, zeroAddress } from 'viem'
import { baseSepolia, base } from 'viem/chains'
import foundnoneVrfAbi from './abi/foundnone-vrf.json'
import { wordlists } from 'bip39'

export default function Home() {
  const IS_PROD = process.env.NEXT_PUBLIC_ENV === 'production'
  const CONTRACT_ADDRESS: `0x${string}` = process.env.NEXT_PUBLIC_VRF_CONTRACT_ADDRESS as `0x${string}` || '0x6011C31271b321FcE089FB898ecd487BA96CC73f'
  const RPC_URL = process.env.NEXT_PUBLIC_RPC_URL
  const [client, setClient] = useState<WalletClient | null>(null)
  const [account, setAccount] = useState<`0x${string}`>()
  const [rand, setRand] = useState<string | null>(null)
  const [contractBalance, setContractBalance] = useState<string | null>(null)
  const [fulfillerBalance, setFulfillerBalance] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const [initializing, setInitializing] = useState(true)
  const [terminalInput, setTerminalInput] = useState<string>('')
  const [terminalOutput, setTerminalOutput] = useState<string[]>([])

  const inputRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    const terminal = document.getElementById('terminal')
    if (terminal) {
      terminal.scrollTop = terminal.scrollHeight
    }
  }, [terminalOutput])

  useEffect(() => {
    const handleKeyPress = (e: KeyboardEvent) => {
      if (e.key === 'c' || e.key === 'r' || e.key === 'b' || e.key === 'l') {
        e.preventDefault()
        setTerminalInput(prev => prev + e.key)

        if (inputRef.current) {
          inputRef.current.focus()
        }
      }
    }
    document.addEventListener('keypress', handleKeyPress)
    return () => {
      document.removeEventListener('keypress', handleKeyPress)
    }
  }, [])



  const LoadingDots = (
    <span className="animate-pulse">
      <span className="inline-block w-2 h-2 bg-green-400 rounded-full mr-1"></span>
      <span className="inline-block w-2 h-2 bg-green-400 rounded-full mr-1"></span>
      <span className="inline-block w-2 h-2 bg-green-400 rounded-full"></span>
    </span>
  )

  const publicClient = createPublicClient({
    chain: IS_PROD ? base : baseSepolia,
    transport: http(
      RPC_URL ? RPC_URL : ''
    ),
  })

  useEffect(() => {
    if (client) {
      updateBalances()
    }
  }, [client, account, rand])

  async function updateBalances() {
    try {
      const contractFullBalance = await publicClient.getBalance({
        address: CONTRACT_ADDRESS,
      })
      const contractFeeBalance = await publicClient.readContract({
        address: CONTRACT_ADDRESS,
        abi: foundnoneVrfAbi,
        functionName: 'contractFeeBalance',
      }) as any
      const aggregateFulfillerBalance = BigInt(contractFullBalance.toString()) - BigInt(contractFeeBalance.toString())
      setContractBalance(contractFeeBalance.toString())
      setFulfillerBalance(aggregateFulfillerBalance.toString())
    } catch (e) {
      console.error('Error updating balances:', e)
    } finally {
      setInitializing(false)
    }
  }

  async function connectWallet() {
    const provider = (window as any).ethereum
    if (!provider) {
      appendTerminalOutput('No wallet provider found.')
      return
    }

    if (account) {
      await provider.request({
        method: 'wallet_requestPermissions',
        params: [{ eth_accounts: {} }],
      })
    }
    try {
      await provider.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: `0x${IS_PROD ? base.id.toString(16) : baseSepolia.id.toString(16)}` }],
      })
    } catch (error: any) {
      if (error.code === 4902) {
        await provider.request({
          method: 'wallet_addEthereumChain',
          params: [{
            chainId: `0x${IS_PROD ? base.id.toString(16) : baseSepolia.id.toString(16)}`,
            chainName: IS_PROD ? 'BASE Mainnet' : 'BASE Sepolia Testnet',
            rpcUrls: [(IS_PROD ? 'https://base.gateway.tenderly.co' : 'https://base-sepolia.gateway.tenderly.co')],
            blockExplorerUrls: [IS_PROD ? 'https://basescan.org' : 'https://sepolia.basescan.org'],
            nativeCurrency: {
              name: IS_PROD ? 'Base' : 'Base Sepolia',
              symbol: IS_PROD ? 'ETH' : 'ETH',
              decimals: 18,
            },
          }],
        })
      } else if (error.code === 4001) {
        appendTerminalOutput('User rejected the request to switch networks.')
        return
      } else {
        console.error('Error switching network:', error)
        appendTerminalOutput('Error switching network. Make sure you are connected to' + (IS_PROD ? ' BASE mainnet.' : ' BASE Sepolia testnet.'))
        return
      }
    }

    const accounts: string[] = await provider.request({ method: 'eth_requestAccounts' })
    const acc = accounts[0] as `0x${string}`
    setAccount(acc)

    const walletClient = createWalletClient({
      transport: custom((window as any).ethereum),
      chain: IS_PROD ? base : baseSepolia,
      account: acc,
    })
    setClient(walletClient)
    appendTerminalOutput(`Connected wallet: ${acc}`)
  }

  async function requestRng() {
    if (!client || !account) {
      appendTerminalOutput('Connect wallet first.')
      return
    }
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
        chain: IS_PROD ? base : baseSepolia,
        abi: foundnoneVrfAbi,
        functionName: 'requestRng',
        value: vrfFee as any,
        args: [zeroAddress, 0]
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
      appendTerminalOutput(`RNG requested. Request ID: ${requestId}`)
      await pullForEntropy(requestId)
    } catch (error) {
      console.error('Error requesting RNG:', error)
      appendTerminalOutput('Error requesting RNG.')
    } finally {
      setLoading(false)
    }
  }

  async function pullForEntropy(requestId: string) {
    await new Promise(resolve => setTimeout(resolve, 3000))
    let tries = 0
    while (tries < 10) {
      const rng = await publicClient.readContract({
        address: CONTRACT_ADDRESS,
        abi: foundnoneVrfAbi,
        functionName: 'entropies',
        args: [requestId]
      })
      if (rng) {
        setRand(rng.toString())
        appendTerminalOutput(`Entropy received: ${rng.toString()}`)
        appendTerminalOutput(`Mod 100,000 / 100,000: ${(parseInt(rng.toString()) % 100000) / 100000}`)
        appendTerminalOutput(randomWords(rng.toString()))
        break
      }
      tries++
      await new Promise(resolve => setTimeout(resolve, 1000))
    }
  }

  const randomWords = (entropy: string) => {
    if (!entropy) return 'Error generating random words.'
    const WORD_COUNT = 8
    const hex = BigInt(entropy).toString(16)
    const bits = hex.padStart(WORD_COUNT * 4, '0')
    const chunkSize = bits.length / WORD_COUNT
    const w = Array.from({ length: WORD_COUNT }, (_, i) => {
      const slice = bits.slice(i * chunkSize, (i + 1) * chunkSize)
      const idx = parseInt(slice, 16) % wordlists.english.length
      return wordlists.english[idx]
    })
    return `Random words: ${w.join(' ')}`
  }

  function appendTerminalOutput(line: string) {
    setTerminalOutput(prev => [...prev, line])
  }

  async function handleTerminalCommand(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    const input = terminalInput.trim().toLowerCase()

    appendTerminalOutput(`> ${input}`)

    if (input === 'connect') {
      connectWallet()
    } else if (input === 'rng' || input === 'request rng') {
      requestRng()
    } else if (input === 'balance') {
      await updateBalances()
      if (contractBalance && fulfillerBalance) {
        appendTerminalOutput(`Contract balance: ${contractBalance} wei`)
        appendTerminalOutput(`Fulfiller balance: ${fulfillerBalance} wei`)
      }
    } else if (input === 'clear') {
      setTerminalOutput([])
    } else {
      appendTerminalOutput('Use "connect" to connect your wallet, "rng" to request a random number.')
    }
    setTerminalInput('')
  }

  return (
    <div className="flex items-center justify-center h-screen w-full bg-[#14101e] no-scrollbar break-all">
      <div className="bg-black text-green-400 font-mono p-6 rounded-lg shadow-inner w-auto max-w-[1200px] h-[600px] flex flex-col no-scrollbar">
        <div className="mb-2 flex flex-col items-start justify-center gap-2">
          <h1 className="text-2xl font-bold mb-2">Foundnone VRF</h1>
          <p className="text-sm">A democratized VRF allowing anyone to request and fulfill entropy requests onchain for rewards on Ethereum.</p>
          <p className="text-sm">
            {IS_PROD ? 'This is a production implementation on BASE with contract address:' : 'This is a test implementation on BASE SEPOLIA with contract address:'}
            <a href={`https://${IS_PROD ? '' : 'sepolia.'}basescan.org/address/${CONTRACT_ADDRESS}`} target="_blank" rel="noreferrer" className='text-sm underline inline-block ml-1'>{CONTRACT_ADDRESS}</a>
          </p>
          <p className="text-sm">
            Codebase:
            <a href={`https://github.com/transmental/foundnone-vrf`} target="_blank" rel="noreferrer" className='text-sm underline inline-block ml-1'>{`https://github.com/transmental/foundnone-vrf`}</a>
          </p>
          <p className="text-sm">
            Connect with me:
            <a href={`https://x.com/transmental`} target="_blank" rel="noreferrer" className='text-sm underline inline-block ml-1'>{`https://x.com/transmental`}</a>
          </p>
          <p className="text-sm">
            Join the Foundnone Research discord:
            <a href={`https://discord.gg/5EGw6szwQp`} target="_blank" rel="noreferrer" className='text-sm underline inline-block ml-1'>{`https://discord.gg/5EGw6szwQp`}</a>
          </p>
        </div>
        <div className="flex-1 overflow-y-auto space-y-1 no-scrollbar scroll-auto border-t-1 border-x-1 px-1 border-green-400" id="terminal">
          {terminalOutput.map((line, idx) => (
            <div key={idx}>{line}</div>
          ))}


          {initializing && (
            <div>&gt; Type {'connect'} to get started {LoadingDots}</div>
          )}

          {loading && (
            <div>&gt; {LoadingDots}</div>
          )}
        </div>

        <form onSubmit={handleTerminalCommand} className="flex border-b-1 border-x-1 px-1 border-green-400">
          <span className="mr-2">&gt;</span>
          <input
            ref={inputRef}
            type="text"
            className="flex-1 bg-black text-green-400 outline-none"
            value={terminalInput}
            onChange={(e) => setTerminalInput(e.target.value)}
            autoFocus
          />

        </form>
        <p className="text-sm mt-2">Type `connect` to connect or switch wallets, `rng` to request a random number, or `balance` to check contract balances.</p>
      </div>
    </div>
  )
}
