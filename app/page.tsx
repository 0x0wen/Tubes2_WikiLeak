import Form from './Form'
const Page = () => {
	return (
		<main className="flex justify-center items-center h-full w-full min-h-screen">
			<div className="mx-auto my-auto w-[40rem] h-fit text-center space-y-6">
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
