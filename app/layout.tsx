import '@/styles/globals.css'
import {Inter as FontSans} from 'next/font/google'
import {cn} from '@/lib/utils'
import {LayoutProps} from '@/.next/types/app/layout'
import {Toaster} from '@/components/ui/toaster'

const fontSans = FontSans({
	subsets: ['latin'],
	variable: '--font-sans',
})

export default function RootLayout({children}: LayoutProps) {
	return (
		<html lang="en" suppressHydrationWarning>
			<head />
			<body
				className={cn(
					'min-h-screen bg-primary text-primary-foreground font-sans antialiased ',
					fontSans.variable
				)}
			>
				{children}
				<Toaster />
			</body>
		</html>
	)
}
