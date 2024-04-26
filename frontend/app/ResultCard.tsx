'use client'
// @ts-ignore
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from '@/components/ui/card'
import {PathType, useResultStore} from '@/store/store'
import pikachu from '@/public/assets/pikachu.gif'
import Image from 'next/image'
import {Separator} from '@/components/ui/separator'
import {isRegularPath} from './functions'
const ResultCard = () => {
	const {isLoading, result, setIsLoading} = useResultStore()
	return (
		<>
			{isLoading ? (
				<div className="w-full min-h-screen flex flex-col justify-center items-center">
					<Image src={pikachu} alt="pikachu" className="aspect-square w-80" />
					<p className="font-mono pt-2">Finding the leak🔎...</p>
				</div>
			) : (
				result &&
				(isRegularPath(result.path[0]) ? (
					<div className="min-h-screen flex justify-center items-center">
						<Card className="w-[350px] bg-primary text-primary-foreground text-center ">
							<CardHeader>
								<CardTitle>💥🚨LEAKED🚨💥</CardTitle>
								<CardDescription>Don't tell the feds!</CardDescription>
							</CardHeader>
							<CardContent>
								<div className="mb-6 w-full overflow-y-auto">
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
								<div className="grid w-full items-center gap-4 ">
									<div className="flex flex-col  bg-foreground rounded-xl border border-white ">
										{result?.path.map((path, index) => (
											<div key={index}>
												<div className=" grid grid-cols-[25px_1fr] items-center p-4  text-primary-foreground ">
													<p>
														{index === 0
															? '1️⃣'
															: index === 1
															? ' 2️⃣'
															: index === 2
															? '3️⃣'
															: index === 3
															? '4️⃣'
															: index === 4
															? '5️⃣'
															: index === 5
															? '6️⃣'
															: index === 6
															? '7️⃣'
															: index === 7
															? '8️⃣'
															: index === 8
															? '9️⃣'
															: '🔟'}
													</p>
													<a
														href={`https://en.wikipedia.org${
															(path as PathType).Link
														}`}
														target="_blank"
														className="text-primary-foreground hover:underline text-sm w-full font-medium leading-none"
													>
														{(path as PathType).Title
															? (path as PathType).Title
															: (path as PathType).Link.replace(
																	'/wiki/',
																	''
															  ).replace(/_/g, ' ')}
													</a>
												</div>
												{index + 1 !== result?.path.length && <Separator />}
											</div>
										))}
									</div>
								</div>
							</CardContent>
						</Card>
					</div>
				) : (
					<>
						<Card className="w-[350px] bg-primary text-primary-foreground text-center mt-40">
							<CardHeader>
								<CardTitle>💥🚨LEAKED🚨💥</CardTitle>
								<CardDescription>Don't tell the feds!</CardDescription>
							</CardHeader>
							<CardContent>
								<div className=" w-full overflow-y-auto">
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
							</CardContent>
						</Card>
						<div className="grid grid-cols-3 py-10 gap-10">
							{result?.path.map((solutions, index) => (
								<div className="flex flex-col rounded-xl border border-white ">
									{(solutions as PathType[]).map((solution, index) => (
										<div key={index}>
											<div className=" grid grid-cols-[25px_1fr] items-center p-4  text-primary-foreground ">
												<p>
													{index === 0
														? '1️⃣'
														: index === 1
														? ' 2️⃣'
														: index === 2
														? '3️⃣'
														: index === 3
														? '4️⃣'
														: index === 4
														? '5️⃣'
														: index === 5
														? '6️⃣'
														: index === 6
														? '7️⃣'
														: index === 7
														? '8️⃣'
														: index === 8
														? '9️⃣'
														: '🔟'}
												</p>
												<a
													href={`https://en.wikipedia.org${solution.Link}`}
													target="_blank"
													className="text-primary-foreground hover:underline text-sm w-full font-medium leading-none"
												>
													{solution.Title
														? solution.Title
														: solution.Link.replace('/wiki/', '').replace(
																/_/g,
																' '
														  )}
												</a>
											</div>
											{index + 1 !== result?.path.length && <Separator />}
										</div>
									))}
								</div>
							))}
						</div>
					</>
				))
			)}
		</>
	)
}

export default ResultCard
