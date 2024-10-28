import HeaderLogo from "../components/headers/HeaderLogo"

function Transaction() {
  return (
    <Scroll>
      <HeaderLogo />
      <div className="mb-4 mt-6 px-6">
        <div className="mb-4 flex items-center justify-between">
          <h1 className="text-2xl font-bold">Transaction</h1>
        </div>
        <div className="mb-4 flex justify-between text-sm">
          <p>Transport 10 Km from lyf</p>
          <p>15$</p>
        </div>
        <div className="flex flex-col gap-2">
          <h5 className="text-sm font-bold">In Destination</h5>
          <div className="flex items-center rounded p-3 shadow-md">
            <div className="flex flex-grow flex-col gap-1">
              <h5 className="text-sm font-bold">Tourist Attraction</h5>
              <p>Premium Ticket</p>
            </div>
            <p className="text-sm">10$</p>
          </div>
          <div className="flex items-center rounded p-3 shadow-md">
            <div className="flex flex-grow flex-col gap-1">
              <h5 className="text-sm font-bold">Tourist Attraction</h5>
              <p>Premium Ticket</p>
            </div>
            <p className="text-sm">10$</p>
          </div>
        </div>
      </div>
      <div className="absolute bottom-0 h-auto w-96 border-t bg-white">
        <div className="flex h-full w-full flex-col gap-3 px-6 pb-4 pt-2">
          <div className="flex justify-between">
            <h1 className="text-xl font-bold">Total</h1>
            <h1 className="text-xl font-bold">15$</h1>
          </div>
          <button className="h-fit w-auto rounded bg-yellow px-4 py-2 text-sm shadow-md">Next</button>
        </div>
      </div>
    </Scroll>
  )
}

export default Transaction