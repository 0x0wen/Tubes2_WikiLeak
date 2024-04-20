import Form from './Form'
import ParticlesBackground from './ParticlesBackground'
import ResultCard from './ResultCard'
const Page = () => {
	return (
		<>
			<main className="w-full h-screen flex flex-col">
				<ParticlesBackground />
				<section className="w-full h-screen flex ">
					<div className="mx-auto my-auto w-full min-h-screen text-center space-y-6 flex justify-center items-center gap-10 z-50">
						<section className="text-left w-fit">
							<h1 className="font-mechanized text-8xl ">WikiLeak.</h1>
							<p className="font-mono text-2xl">
								<span className="font-semibold">Leak</span> the{' '}
								<span className="strike">
									<s>deepest</s>
								</span>{' '}
								<span className="underline">connection</span> of{' '}
								<span className="text-red-700">things</span>.
							</p>
						</section>
						<Form />
					</div>
				</section>
				<div className="w-full min-h-screen z-50 flex justify-center items-center">
					<ResultCard />
				</div>
			</main>
		</>
	)
}

export default Page
