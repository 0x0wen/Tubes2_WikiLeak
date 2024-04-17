'use client'
import {Input} from '../components/ui/input'
import {RadioGroup, RadioGroupItem} from '../components/ui/radio-group'
import {Label} from '../components/ui/label'
import {Button} from '../components/ui/button'
import {Loader2} from 'lucide-react'
import {ChangeEvent, FormEvent, useState} from 'react'
import {Skeleton} from '../components/ui/skeleton'
import {useToast} from '../components/ui/use-toast'
import axios from 'axios'
import get from 'lodash/get'
import filter from 'lodash/filter'
import forEach from 'lodash/forEach'
import debounce from 'lodash/debounce'
import {WIKIPEDIA_API_URL} from './constants'
import Image from 'next/image'
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
interface suggestionType {
	title: string
	description: string
	thumbnailUrl: string
}

const Form = () => {
	const [isLoading, setIsLoading] = useState<boolean>(false)
	const [result, setResult] = useState<{time: number} | undefined>()
	const [startSuggestions, setStartSuggestions] = useState<suggestionType[]>([])
	const [openStartSuggestions, setStartOpenSuggestions] =
		useState<boolean>(false)
	const [goalSuggestions, setGoalSuggestions] = useState<suggestionType[]>([])
	const [openGoalSuggestions, setGoalOpenSuggestions] = useState<boolean>(false)
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
	const onInputChange = async (e: ChangeEvent<HTMLInputElement>) => {
		const value = e.target.value
		const id = e.target.id
		if (id === 'start') {
			setConfig({...config, start: e.target.value})
		} else {
			setConfig({...config, goal: e.target.value})
		}
		const queryParams = {
			action: 'query',
			format: 'json',
			gpssearch: value,
			generator: 'prefixsearch',
			prop: 'pageprops|pageimages|pageterms',
			redirects: '', // Automatically resolve redirects
			ppprop: 'displaytitle',
			piprop: 'thumbnail',
			pithumbsize: '160',
			pilimit: '30',
			wbptterms: 'description',
			gpsnamespace: 0, // Only return results in Wikipedia's main namespace
			gpslimit: 5, // Return at most five results
			origin: '*',
		}
		await axios({
			method: 'get',
			url: WIKIPEDIA_API_URL,
			params: queryParams,
		})
			.then(async (response) => {
				const suggestion: suggestionType[] = []
				const pageResults = get(response, 'data.query.pages', {})
				await forEach(
					pageResults,
					({
						ns,
						index,
						title,
						terms,
						thumbnail,
					}: {
						ns: any
						index: any
						title: any
						terms: any
						thumbnail: any
					}) => {
						// Due to https://phabricator.wikimedia.org/T189139, results will not always be limited
						// to the main namespace (0), so ignore all results which have a different namespace.
						if (ns === 0) {
							let description = get(terms, 'description.0')
							if (description) {
								description =
									description.charAt(0).toUpperCase() + description.slice(1)
							}

							suggestion[index - 1] = {
								title,
								description,
								thumbnailUrl: get(thumbnail, 'source'),
							}
						}
					}
				)
				// Due to ignoring non-main namespaces above, the suggestions array may have some missing
				// items, so remove them via filter().
				if (id == 'start') {
					setStartSuggestions(filter(suggestion))
					setStartOpenSuggestions(true)
				} else {
					setGoalSuggestions(filter(suggestion))
					setGoalOpenSuggestions(true)
				}
			})
			.catch((error) => {
				// Report the error to Google Analytics, but don't report any user-facing error since the
				// input is still usable even without suggestions.
				const defaultErrorMessage =
					'Request to fetch page suggestions from Wikipedia API failed.'
				console.log(error)
				// window.ga('send', 'exception', {
				// 	exDescription: get(error, 'response.data.error', defaultErrorMessage),
				// 	exFatal: false,
				// })
			})
		console.log(startSuggestions)
	}
	return (
		<div className="w-max mx-auto my-auto ">
			<form
				className=" space-y-4 w-full flex flex-col justify-center items-center"
				onSubmit={submitHandler}
				onReset={onClear}
			>
				<div className="grid w-80 max-w-sm items-center gap-1.5 text-left relative">
					<Label htmlFor="start">Start Title</Label>
					<Input
						type="text"
						id="start"
						placeholder="Title..."
						value={config.start}
						onChange={onInputChange}
						onFocus={() => setGoalOpenSuggestions(false)}
					/>
					{openStartSuggestions && (
						<div className="absolute top-16 w-full z-10">
							{startSuggestions.length > 0 ? (
								startSuggestions.map((suggestion, index) => (
									<button
										key={index}
										className="flex items-center space-x-2 p-2 text-primary bg-primary-foreground hover:bg-slate-200 w-full text-left"
										onClick={() => {
											setConfig({...config, start: suggestion.title})
											setStartOpenSuggestions(false)
										}}
										type="button"
									>
										<Image
											src={suggestion.thumbnailUrl}
											alt={suggestion.title}
											className="w-8 h-8  object-cover"
											width={50}
											height={50}
										/>
										<div>
											<label>{suggestion.title}</label>
											<p className="text-sm text-gray-500">
												{suggestion.description}
											</p>
										</div>
									</button>
								))
							) : (
								<div className="flex items-center justify-center space-x-2 p-2 text-primary bg-primary-foreground hover:bg-slate-200 w-full text-left">
									No available page
								</div>
							)}
						</div>
					)}
				</div>

				<div className="grid w-80 max-w-sm items-center gap-1.5 text-right relative">
					<Label htmlFor="goal">Goal Title</Label>
					<Input
						type="text"
						id="goal"
						placeholder="Title..."
						value={config.goal}
						className="text-right placeholder:text-right"
						onChange={onInputChange}
						onFocus={() => setStartOpenSuggestions(false)}
					/>
					{openGoalSuggestions && (
						<div className="absolute top-16 w-full z-10">
							{goalSuggestions.length > 0 ? (
								goalSuggestions.map((suggestion, index) => (
									<button
										key={index}
										className="flex items-center space-x-2 p-2 text-primary bg-primary-foreground hover:bg-slate-200 w-full text-left"
										onClick={() => {
											setConfig({...config, goal: suggestion.title})
											setGoalOpenSuggestions(false)
										}}
										type="button"
									>
										<Image
											src={suggestion.thumbnailUrl}
											alt={suggestion.title}
											className="w-8 h-8 object-cover"
											width={50}
											height={50}
										/>
										<div>
											<label>{suggestion.title}</label>
											<p className="text-sm text-gray-500">
												{suggestion.description}
											</p>
										</div>
									</button>
								))
							) : (
								<div className="flex items-center justify-center space-x-2 p-2 text-primary bg-primary-foreground hover:bg-slate-200 w-full">
									No available page
								</div>
							)}
						</div>
					)}
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
