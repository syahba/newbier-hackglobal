import { useNavigate } from "react-router-dom";
import ButtonAction from "../components/buttons/ButtonAction";
import HeaderLogo from "../components/headers/HeaderLogo"
import CardProduct from "../components/itineraries/CardProduct"
import Scroll from "../layouts/Scroll"

function Market() {
  const navigate = useNavigate();

  const data = {
    name: 'Lyf Funan',
    star: '4.5'
  };
  
  return (
    <Scroll>
      <HeaderLogo />
      
      <div className="px-6 mt-2 mb-4">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold">Itinerary</h1>
        </div>
        <div className="flex flex-col gap-2">
          <h5 className="text-base font-bold">In Destination</h5>
          <CardProduct data={data} />
        </div>
      </div>

      <div className="absolute bottom-0 h-auto bg-white border-t w-96">
        <div className="flex items-center justify-between w-full h-full px-6 py-4">
          <div>
            <h1 className="text-xl font-bold">15$</h1>
          </div>

          <ButtonAction text={'Next'} onClick={() => navigate('/transaction')} />
        </div>
      </div>
    </Scroll>
  )
}

export default Market