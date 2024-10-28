import ButtonAction from "../components/buttons/ButtonAction"
import HeaderDefault from "../components/headers/HeaderDeafult"
import Itinerary from "../components/itineraries/Itinerary"
import Scroll from "../layouts/Scroll"

function HomeItinerary () {

  return (
    <Scroll>
      <HeaderDefault />

      <div className="flex flex-col gap-4 px-6 mt-2 mb-4">
        <div>
          <h1 className="mb-2 text-2xl font-bold">Travel Buddy</h1>
          <p className="">Waiting for a match...</p>
          <div className="relative hidden rounded shadow-xl bg-blue">
            <div className="flex flex-col gap-2 px-3 pt-4 pb-6 text-sm text-white">
              <h5 className="font-bold">User2 would like to join your travel!</h5>
              <div className="flex gap-8">
                <p>Thailand</p>
                <p>23</p>
              </div>
              <p>I love thrilling roller coasters and spooky attractions! I’m excited to conquer every ride at the USS and experience the chills at the haunted maze!</p>
            </div>
            <div className="absolute -bottom-5 right-4">
              <div className="flex gap-3">
                <button className="rounded-full h-11 w-11 bg-yellow"></button>
                <button className="bg-white border-2 rounded-full h-11 w-11 border-yellow"></button>
              </div>
            </div>
          </div>
          <div className="relative hidden rounded shadow-xl bg-blue">
            <div className="flex flex-col gap-2 px-3 pt-4 pb-6 text-sm text-white">
              <div className="flex gap-8">
                <p>User2</p>
                <p>Thailand</p>
                <p>23</p>
              </div>
              <p>I love thrilling roller coasters and spooky attractions! I’m excited to conquer every ride at the USS and experience the chills at the haunted maze!</p>
            </div>
            <div className="absolute -bottom-5 right-4">
              <div className="flex gap-3">
                <button className="rounded-full h-11 w-11 bg-yellow"></button>
              </div>
            </div>
          </div>
        </div>
        <div>
          <h1 className="mb-2 text-2xl font-bold">Product</h1>
          <a href="">
            <div className="px-4 py-2 border rounded w-fit border-blue">
              <p>See in market</p>
            </div>
          </a>
        </div>
        
        <div>
          <h1 className="mb-2 text-2xl font-bold">Itinerary</h1>

          <Itinerary data={[]} />
        </div>

        <div className="flex items-center justify-center w-full h-full py-4">
          <ButtonAction text="Completed" />
        </div>
      </div>
    </Scroll>
  )
}

export default HomeItinerary