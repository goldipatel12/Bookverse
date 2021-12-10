import '../styles/globals.css'
import Link from 'next/link'

function Marketplace({ Component, pageProps }) {
  return (
    <div>
      <nav className="border-b p-6">
        <div className="flex mt-4">
        <p className="text-5xl text-blue-500 font-bold">Bookverse Marketplace</p>
        <span className="text-2xl">Beta</span>
        </div>
        <div className="flex mt-4">
          <Link href="/">
            <a className="mr-6 text-blue-1000">
              Home
            </a>
          </Link>
          <Link href="/create-item">
            <a className="mr-6 text-blue-1000">
              Sell NFT Books
            </a>
          </Link>
          <Link href="/my-assets">
            <a className="mr-6 text-blue-1000">
              My Books
            </a>
          </Link>
          <Link href="/creator-dashboard">
            <a className="mr-6 text-blue-1000">
              Author's Dashboard
            </a>
          </Link>
        </div>
      </nav>
      <Component {...pageProps} />
    </div>
  )
}

export default Marketplace