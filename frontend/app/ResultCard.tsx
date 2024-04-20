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
import {useResultStore} from '@/store/store'
const ResultCard = () => {
	const {isLoading, result, setIsLoading} = useResultStore()
	return (
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
									40s
								</td>
							</tr>
							<tr className="m-0 border-t p-0 ">
								<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
									Checked article(s)
								</td>
								<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
									323
								</td>
							</tr>
							<tr className="m-0 border-t p-0 ">
								<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
									Passed article(s)
								</td>
								<td className="border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right">
									32
								</td>
							</tr>
						</tbody>
					</table>
				</div>
				<div className="grid w-full items-center gap-4">
					<div className="flex flex-col space-y-1.5">
						<Label htmlFor="name">Name</Label>
						<p>335ms</p>
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
}

export default ResultCard
