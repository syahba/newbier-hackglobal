function Home() {
  return (
    <>
    <div className="px-6 mt-12 mb-4">
      <h1 className="text-2xl font-bold">Hi, <span className="text-blue">User</span>!</h1>
      <h5 className="mb-4 text-sm font-bold">Where would you like to go today?</h5>
      <input type="text" className="mb-4 h-8 w-full px-3 shadow-md" placeholder="Enter a destination" />
      <div className="flex w-full gap-4">
        <button className="rounded border border-blue px-5 py-2 shadow-md">I’m not sure</button>
        <button className="rounded border flex-auto border-yellow px-5 py-2 shadow-md">I’d like to travel with a buddy</button>
      </div>
    </div>
    <div className="absolute bottom-0 w-96 h-[449px]">
      <div className="bg-home h-full w-96 bg-contain bg-center bg-no-repeat"></div>
    </div>
  </>
  )
}

export default Home