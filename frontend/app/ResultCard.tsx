'use client'
import {Button} from '@/components/ui/button'
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from '@/components/ui/card'
import {Label} from '@/components/ui/label'
import {TypographyP} from '@/components/ui/typography'
import {useResultStore} from '@/store/store'
import pikachu from '@/public/assets/pikachu.gif'
import Image from 'next/image'
const ResultCard = () => {
	const {isLoading, result, setIsLoading} = useResultStore()
	return (
		<>
			{isLoading ? (
				<div className="flex flex-col justify-center items-center">
					<Image src={pikachu} alt="pikachu" className="aspect-square w-80" />
					<p className="font-mono pt-2">Finding the leak🔎...</p>
				</div>
			) : (
				result && (
					<Card className="w-[350px] bg-primary text-primary-foreground text-center">
						<CardHeader>
							<CardTitle>💥🚨LEAKED🚨💥</CardTitle>
							<CardDescription>Don't tell the feds!</CardDescription>
						</CardHeader>
						<CardContent>
							<div className="my-6 w-full overflow-y-auto">
								<table className="w-full">
									<tbody>
										<tr className="m-0 border-t p-0 ">
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												Time
											</td>
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												{result?.time.toFixed(2)} ms
											</td>
										</tr>
										<tr className="m-0 border-t p-0 ">
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												Checked articles
											</td>
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												{result?.checkedArticles}
											</td>
										</tr>
										<tr className="m-0 border-t p-0 ">
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												Passed articles
											</td>
											<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
												{result?.passedArticles}
											</td>
										</tr>
									</tbody>
								</table>
							</div>
							<div className="grid w-full items-center gap-4">
								<div className="flex flex-col space-y-1.5">
									{result?.path.map((path, index) => (
										<div
											key={index}
											className="mb-4 grid grid-cols-[25px_1fr] items-start pb-4 last:mb-0 last:pb-0 text-primary-foreground"
										>
											<span className="flex h-2 w-2 translate-y-1 rounded-full bg-sky-500" />
											<div className="space-y-1">
												<p className="text-sm font-medium leading-none">
													{path.Title}
												</p>
												<p className="text-sm text-muted-foreground">
													{path.Link}
												</p>
											</div>
										</div>
									))}
								</div>
								<div className="flex flex-col space-y-1.5"></div>
							</div>
						</CardContent>
						<CardFooter className="flex justify-between">
							<Button variant="outline">Cancel</Button>
							<Button>Deploy</Button>
						</CardFooter>
					</Card>
				)
			)}
		</>
	)
}

export default ResultCard
