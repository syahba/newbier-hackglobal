import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"

function Market() {
  return (
    <Scroll>
      <HeaderLogo />
      
      <div className="px-6 mt-6 mb-4">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold">Itinerary</h1>
        </div>
        <div className="flex justify-between p-3 mb-4 text-white rounded bg-blue">
          <label for="transport-destination" className="flex flex-grow flex-col gap-0.5">
            <h5 className="text-sm font-bold">Transport to first destination</h5>
            <p>5$ per Km from lyf hotel</p>
            <p className="text-[0.6rem]">*) Cost will be split if you travel with a buddy</p>
          </label>
          <input type="checkbox" name="destination-transport" id="transport-destination" className="self-center w-5 h-5" />
        </div>
        <div className="flex flex-col gap-2">
          <h5 className="text-sm font-bold">In Destination</h5>
          <div className="rounded shadow-md">
            <div className="flex flex-grow gap-3 p-3">
              <div className="w-16 h-16 rounded bg-blue"></div>
              <div className="flex flex-col flex-grow gap-1">
                <h5 className="text-sm font-bold">Lyf Funan</h5>
                <div className="flex gap-2">
                  <div className="inline-block w-4 h-4 bg-yellow"></div>
                  4.5
                </div>
                <div className="flex gap-2">
                  <div className="inline-block w-4 h-4 bg-yellow"></div>
                  Singapore
                </div>
              </div>
              <div className="self-center">
                <svg className="ms-3 h-2.5 w-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 4 4 4-4" />
                </svg>
              </div>
            </div>
            <div id="dropdown" className="flex flex-col gap-2 p-3">
              <div className="flex items-center gap-2 text-sm">
                <p className="flex-grow">Premium Ticket</p>
                <p>10$/item</p>
                <input type="text" className="px-1 py-0 text-center border rounded w-7 border-blue" placeholder="0" />
              </div>
              <div className="flex items-center gap-2 text-sm">
                <p className="flex-grow">Premium Ticket</p>
                <p>10$/item</p>
                <input type="text" className="px-1 py-0 text-center border rounded w-7 border-blue" placeholder="0" />
              </div>
            </div>
          </div>
          <div className="rounded shadow-md">
            <div className="flex flex-grow gap-3 p-3">
              <div className="w-16 h-16 rounded bg-blue"></div>
              <div className="flex flex-col flex-grow gap-1">
                <h5 className="text-sm font-bold">Lyf Funan</h5>
                <div className="flex gap-2">
                  <div className="inline-block w-4 h-4 bg-yellow"></div>
                  4.5
                </div>
                <div className="flex gap-2">
                  <div className="inline-block w-4 h-4 bg-yellow"></div>
                  Singapore
                </div>
              </div>
              <div className="self-center">
                <svg className="ms-3 h-2.5 w-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 4 4 4-4" />
                </svg>
              </div>
            </div>
            <div className="flex flex-col hidden gap-2 p-3">
              <div className="flex items-center gap-2 text-sm">
                <p className="flex-grow">Premium Ticket</p>
                <p>10$/item</p>
                <input type="text" className="px-1 py-0 text-center border rounded w-7 border-blue" placeholder="0" />
              </div>
              <div className="flex items-center gap-2 text-sm">
                <p className="flex-grow">Premium Ticket</p>
                <p>10$/item</p>
                <input type="text" className="px-1 py-0 text-center border rounded w-7 border-blue" placeholder="0" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="absolute bottom-0 h-auto bg-white border-t w-96">
        <div className="flex items-center justify-between w-full h-full px-6 py-4">
          <div>
            <h1 className="text-xl font-bold">15$</h1>
            <p className="text-[0.6rem]">*) Transportation is excluded</p>
          </div>
          <button className="w-auto px-4 py-2 text-sm rounded shadow-md h-fit bg-yellow">Next</button>
        </div>
      </div>
    </Scroll>
  )
}

export default Market