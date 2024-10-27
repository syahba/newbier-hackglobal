import ButtonLine from "../components/buttons/ButtonLine"
import ButtonRound from "../components/buttons/ButtonRound"
import CardProfile from "../components/forms/CardProfile"
import HeaderWhite from "../components/headers/HeaderWhite"
import Itinerary from "../components/itineraries/Itinerary"
import Scroll from "../layouts/Scroll"

function Match() {
  return (
    <Scroll>
      <div className="flex flex-col items-center w-full h-96 bg-blue">
        <HeaderWhite />

        <CardProfile />

        <div className="relative flex h-fit w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
          <div className="absolute flex gap-4 -top-6 right-8">
            <ButtonRound />
            
            <ButtonLine />
          </div>

          <div className="my-4">
            <h1 className="mb-2 text-2xl font-bold">Itinerary</h1>

            <Itinerary  data={[]} />
          </div>
        </div>
      </div>
    </Scroll>
  )
}

export default Match