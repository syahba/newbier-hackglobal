import { useNavigate } from "react-router-dom";
import HeaderLogo from "../components/headers/HeaderLogo";
import Scroll from "../layouts/Scroll";
import CardTransaction from "../components/itineraries/CardTransaction";

function Transaction() {
  const navigate = useNavigate();

  const data = [
    {
      id: 1,
      name: "Tourist Attraction",
      star: "4.5",
      product: "Premium Ticket",
      price: "10$/item",
    },
    {
      id: 2,
      name: "Tourist Attraction",
      star: "4.9",
      product: "Gloves",
      price: "5$/item",
    },
  ];

  return (
    <Scroll>
      <HeaderLogo />

      <div className="mb-4 mt-6 px-6">
        <div className="mb-4 flex items-center justify-between">
          <h1 className="text-2xl font-bold">Transaction</h1>
        </div>
        <div className="flex flex-col gap-2">
          <h5 className="text-base font-bold">In Destination</h5>
          {data.map((v) => (
            <div key={v.id}>
              <CardTransaction data={v} />
            </div>
          ))}
        </div>
      </div>
      <div className="absolute bottom-0 h-auto w-96 bg-white shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
        <div className="flex h-full w-full flex-col gap-3 px-6 pb-4 pt-2">
          <div className="flex justify-between">
            <h1 className="text-xl font-bold">Total</h1>
            <h1 className="text-xl font-bold">15$</h1>
          </div>
          <button
            className="h-fit w-auto rounded bg-yellow px-4 py-2 text-sm shadow-md"
            onClick={() => navigate("/home-itinerary")}
          >
            Next
          </button>
        </div>
      </div>
    </Scroll>
  );
}

export default Transaction;
