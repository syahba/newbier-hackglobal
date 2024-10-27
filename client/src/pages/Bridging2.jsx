function Bridging2() {
  return (
    <div className="bg-blue w-full h-screen relative">

      {/* {{include "components/header/header-white"}} */}

      <div className="h-1/5 flex flex-col justify-center w-full px-6 text-center">
        <h1 className="text-xl font-bold text-white">Wonderful! Traveling with a</h1>
        <h1 className="text-xl font-bold text-white mt-1 mb-3">buddy is always a blast!</h1>
        <h5 className="text-sm text-white">Choose your itinerary and get a match!</h5>
      </div>
      <div className="absolute bottom-0 w-96 h-[500px]">
        <div className="bg-bridging2 h-full w-96 bg-contain bg-center bg-no-repeat"></div>
      </div>
      <div className="absolute bottom-5 right-5">Tap to continue</div>
    </div>
  )
}