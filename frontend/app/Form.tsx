'use client'
import {Input} from '../components/ui/input'
import {RadioGroup, RadioGroupItem} from '../components/ui/radio-group'
import {Label} from '../components/ui/label'
import {Button} from '../components/ui/button'
import {Loader2} from 'lucide-react'
import {
	ChangeEvent,
	ChangeEventHandler,
	EventHandler,
	FormEvent,
	useState,
} from 'react'
import {Skeleton} from '../components/ui/skeleton'
import {useToast} from '../components/ui/use-toast'

const Form = () => {
	const [isLoading, setIsLoading] = useState<boolean>(false)
	const [result, setResult] = useState<{time: number} | undefined>()
	const [config, setConfig] = useState<{
		start: string
		goal: string
		algorithm: string
	}>({
		start: '',
		goal: '',
		algorithm: 'IDS',
	})
	const {toast} = useToast()

	const onClear = (e: FormEvent<HTMLFormElement>) => {
		e.preventDefault()
		setConfig({start: '', goal: '', algorithm: 'IDS'})
		toast({title: 'Form cleared', variant: 'default'})
	}
	const submitHandler = (e: FormEvent<HTMLFormElement>) => {
		e.preventDefault()
		if (config.start === '' || config.goal === '') {
			toast({title: 'Please fill all the fields', variant: 'destructive'})
			return
		}
		setIsLoading(true)
		const element = document.getElementById('result')
		setTimeout(() => {
			element?.scrollIntoView({behavior: 'smooth'})
		}, 100)
		console.log(element)
		setTimeout(() => {
			setIsLoading(false)
		}, 2000)
		setResult({time: 2000})
	}
	return (
		<div className="w-max mx-auto my-auto ">
			<form
				className=" space-y-4 w-full flex flex-col justify-center items-center"
				onSubmit={submitHandler}
				onReset={onClear}
			>
				<div className="grid w-80 max-w-sm items-center gap-1.5 text-left">
					<Label htmlFor="start">Start Title</Label>
					<Input
						type="text"
						id="start"
						placeholder="Title..."
						value={config.start}
						onChange={(e) => setConfig({...config, start: e.target.value})}
					/>
				</div>
				<div className="grid w-80 max-w-sm items-center gap-1.5 text-right">
					<Label htmlFor="goal">Goal Title</Label>
					<Input
						type="text"
						id="goal"
						placeholder="Title..."
						value={config.goal}
						className="text-right placeholder:text-right"
						onChange={(e: ChangeEvent<HTMLInputElement>) =>
							setConfig({...config, goal: e.target.value})
						}
					/>
				</div>
				<div>
					<Label htmlFor="algorithm">Algorithm</Label>
					<RadioGroup
						id="algorithm"
						defaultValue="IDS"
						className="flex w-full justify-center py-4"
						onChange={(e: ChangeEvent<HTMLInputElement>) =>
							setConfig({...config, algorithm: e.target.value})
						}
					>
						<div className="flex items-center space-x-2">
							<RadioGroupItem value="IDS" id="r1" />
							<Label htmlFor="r1">IDS</Label>
						</div>
						<div className="flex items-center space-x-2">
							<RadioGroupItem value="BFS" id="r2" />
							<Label htmlFor="r2">BFS</Label>
						</div>
					</RadioGroup>
				</div>
				<div className="flex gap-4 justify-center">
					<Button className="w-32" type="reset">
						Clear
					</Button>
					<Button
						className="w-32"
						variant="secondary"
						disabled={isLoading}
						type="submit"
					>
						{isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
						{isLoading ? 'Loading...' : 'Leak'}
					</Button>
				</div>
			</form>
			<div id="result">
				{isLoading ? (
					<Skeleton className="mt-32 h-96 w-[40rem] rounded-xl mx-auto mb-40" />
				) : (
					result && (
						<div className="mt-32 h-96 w-[40rem] bg-white  mb-40">yessir</div>
					)
				)}
			</div>
		</div>
	)
}

export default Form
