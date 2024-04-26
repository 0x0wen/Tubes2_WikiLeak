'use client'
import {useResultStore} from '@/store/store'
import ResultCard from './ResultCard'
const Result = () => {
	const {result} = useResultStore()
	return (
		<div className="w-full  z-50 flex flex-col justify-center items-center">
			<ResultCard />
		</div>
	)
}

export default Result
