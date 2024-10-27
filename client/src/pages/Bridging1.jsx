import { useNavigate } from "react-router-dom"
import HeaderWhite from "../components/headers/HeaderWhite"
import Main from "../layouts/Main"

function Bridging1() {
  const navigate = useNavigate()
  return (
    <Main>
      <div className="relative w-full h-screen bg-blue">

        <HeaderWhite />

        <div className="flex flex-col justify-center w-full px-6 text-center h-1/5">
          <h5 className="text-sm text-white">Not sure about your destination?</h5>
          <h1 className="mt-3 mb-1 text-xl font-bold text-white">Don’t worry!</h1>
          <h1 className="text-xl font-bold text-white">We’ve got you covered!</h1>
        </div>
        <div className="absolute bottom-0 w-96 h-[563px]">
          <div className="h-full bg-center bg-no-repeat bg-cover bg-bridging w-96"></div>
        </div>
        <div className="absolute cursor-pointer bottom-5 right-5" onClick={() => navigate("/preference")}>Tap to continue</div>
      </div>
    </Main>
  )
}

export default Bridging1