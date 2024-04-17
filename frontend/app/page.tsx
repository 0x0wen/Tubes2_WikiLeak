'use client'
import Form from './Form'

import Image from 'next/image'

const data = {
	name: 'Parent',
	children: [
		{
			name: 'Child One',
			children: [
				{
					name: 'Child One',
				},
				{
					name: 'Child Two',
					children: [
						{
							name: 'Child One',
						},
						{
							name: 'Child Two',
							children: [
								{
									name: 'Child One',
									children: [
										{
											name: 'Child One',
										},
										{
											name: 'Child Two',
										},
									],
								},
								{
									name: 'Child Two',
								},
							],
						},
					],
				},
			],
		},
		{
			name: 'Child Two',
		},
	],
}
const Page = () => {
	return (
		<main className="flex justify-center items-center h-full w-full min-h-screen">
			<div className="mx-auto my-auto w-[40rem] min-h-screen text-center space-y-6 flex flex-col justify-center">
				<section>
					<h1 className="font-mechanized text-7xl ">WikiLeak.</h1>
					<p className="font-mono">Leak the deepest connection of things.</p>
				</section>
				<Form />
			</div>
		</main>
	)
}

export default Page
