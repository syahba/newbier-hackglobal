import { useLocation, useNavigate } from "react-router-dom";
import HeaderWhite from "../components/headers/HeaderWhite";
import Main from "../layouts/Main";

function Bridging3() {
  const navigate = useNavigate();

  const changePage = () => {
    navigate("/home-itinerary");
  };

  return (
    <Main>
      <div className="relative w-full h-screen bg-blue">
        <HeaderWhite />

        <div className="flex flex-col justify-center w-full px-6 text-center h-1/5">
          <h1 className="text-xl font-bold text-white">Awesome! Your travel</h1>
          <h1 className="mt-1 mb-3 text-xl font-bold text-white">
            itinerary is ready.
          </h1>
          <h5 className="mb-2 text-sm text-white">
            Have a excellent trip and stay safe~
          </h5>
        </div>
        <div className="absolute bottom-0 w-96 h-[480px]">
          <div className="h-full bg-center bg-no-repeat bg-contain bg-bridging3 w-96"></div>
        </div>
        <div
          className="absolute bottom-5 right-5 cursor-pointer"
          onClick={changePage}
        >
          Tap to continue
        </div>
      </div>
    </Main>
  );
}

export default Bridging3;
