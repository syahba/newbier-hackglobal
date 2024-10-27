import HeaderLogo from "../components/headers/HeaderLogo"

function Transaction() {
  return (
    <Scroll>
      <HeaderLogo />
      <div class="mb-4 mt-6 px-6">
        <div class="mb-4 flex items-center justify-between">
          <h1 class="text-2xl font-bold">Transaction</h1>
        </div>
        <div class="mb-4 flex justify-between text-sm">
          <p>Transport 10 Km from lyf</p>
          <p>15$</p>
        </div>
        <div class="flex flex-col gap-2">
          <h5 class="text-sm font-bold">In Destination</h5>
          <div class="flex items-center rounded p-3 shadow-md">
            <div class="flex flex-grow flex-col gap-1">
              <h5 class="text-sm font-bold">Tourist Attraction</h5>
              <p>Premium Ticket</p>
            </div>
            <p class="text-sm">10$</p>
          </div>
          <div class="flex items-center rounded p-3 shadow-md">
            <div class="flex flex-grow flex-col gap-1">
              <h5 class="text-sm font-bold">Tourist Attraction</h5>
              <p>Premium Ticket</p>
            </div>
            <p class="text-sm">10$</p>
          </div>
        </div>
      </div>
      <div class="absolute bottom-0 h-auto w-96 border-t bg-white">
        <div class="flex h-full w-full flex-col gap-3 px-6 pb-4 pt-2">
          <div class="flex justify-between">
            <h1 class="text-xl font-bold">Total</h1>
            <h1 class="text-xl font-bold">15$</h1>
          </div>
          <button class="h-fit w-auto rounded bg-yellow px-4 py-2 text-sm shadow-md">Next</button>
        </div>
      </div>
    </Scroll>
  )
}

export default Transaction