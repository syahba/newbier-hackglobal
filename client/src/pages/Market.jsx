import { useNavigate } from "react-router-dom";
import ButtonAction from "../components/buttons/ButtonAction";
import HeaderLogo from "../components/headers/HeaderLogo"
import CardProduct from "../components/itineraries/CardProduct"
import Scroll from "../layouts/Scroll"
import { useEffect } from "react";

function Market() {
  const navigate = useNavigate();

  const data = [{
    id: 1,
    name: 'Tourist Attraction',
    star: '4.5',
    products: [{
      name: 'Premium Ticket',
      price: '10$/item'
    },
    {
      name: 'VIP Ticket',
      price: '20$/item'
    }]
  }, {
    id: 2,
    name: 'Tourist Attraction',
    star: '4.9',
    products: [{
      name: 'Gloves',
      price: '5$/item'
    }]
  }, {
    id: 3,
    name: 'Tourist Attraction',
    star: '5',
    products: [{
      name: 'Surfboard',
      price: '15$/item'
    }]
  }];
  
  return (
    <Scroll>
      <HeaderLogo />
      
      <div className="px-6 mt-2 mb-4">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold">Itinerary</h1>
        </div>
        <div className="flex flex-col gap-2">
          <h5 className="text-base font-bold">In Destination</h5>
          {data.map(v => (
            <div key={v.id}>
              <CardProduct data={v} />
            </div>
          ))}
        </div>
      </div>

        <div className="flex items-center justify-center px-6 py-4">
          <ButtonAction text={'Next'} onClick={() => navigate('/transaction')} />
        </div>
    </Scroll>
  )
}

export default Market