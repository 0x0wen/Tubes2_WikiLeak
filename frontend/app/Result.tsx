'use client'
import {useResultStore} from '@/store/store'
import ResultCard from './ResultCard'
const Result = () => {
	const {result} = useResultStore()
	if (!result) return null
	return (
		<div className="w-full  z-50 flex flex-col justify-center items-center">
			<ResultCard />
		</div>
	)
}

export default Result
