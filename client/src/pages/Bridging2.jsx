import { useNavigate } from "react-router-dom"
import HeaderWhite from "../components/headers/HeaderWhite"
import Main from "../layouts/Main"

function Bridging2() {
  const navigate = useNavigate()
  return (
    <Main>
      <div className="relative w-full h-screen bg-blue">

        <HeaderWhite />

        <div className="flex flex-col justify-center w-full px-6 text-center h-1/5">
          <h1 className="text-xl font-bold text-white">Wonderful! Traveling with a</h1>
          <h1 className="mt-1 mb-3 text-xl font-bold text-white">buddy is always a blast!</h1>
          <h5 className="text-sm text-white">Choose your itinerary and get a match!</h5>
        </div>
        <div className="absolute bottom-0 w-96 h-[500px]">
          <div className="h-full bg-center bg-no-repeat bg-contain bg-bridging2 w-96"></div>
        </div>
        <div className="absolute cursor-pointer bottom-5 right-5" onClick={() => navigate("/preference-buddy")}>Tap to continue</div>
      </div>
    </Main>
  )
}

export default Bridging2