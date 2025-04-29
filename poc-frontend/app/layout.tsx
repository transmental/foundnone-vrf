import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Foundnone VRF",
  description: "A democratized and provably random number generator for Ethereum.",
  openGraph: {
    title: "Foundnone VRF",
    description: "A democratized and provably random number generator for Ethereum.",
    url: "https://vrf.foundnone.xyz",
    images: [
      {
        url: "https://foundnone-home.s3.us-east-1.amazonaws.com/foundnonevrf.png",
        width: 1280,
        height: 640,
      }
    ]
  }
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
