import { useCallback, useState } from "react"
import FieldBuddy from "../components/forms/FieldBuddy"
import FieldDescription from "../components/forms/FieldDescription"
import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"
import ButtonAction from "../components/buttons/ButtonAction"
import { useLocation, useNavigate } from "react-router-dom"

function Buddy() {
  const [isBuddy, setIsBuddy] = useState(false)
  const [description, setDescription] = useState(null)
  const navigate = useNavigate()

  const { state } = useLocation();
  const { itinerary, destination, activity, trip } = state

  const wrapperSetActivity = useCallback(val => {
    setIsBuddy(val);
  }, [setIsBuddy]);
  const wrapperSetDescription = useCallback(val => {
    setDescription(val);
  }, [setDescription]);

  const nextCreate = () => {
    navigate("/buddy/profile", {state: {
      itinerary: itinerary,
      isBuddy: isBuddy,
      description: description,
      destination,
      activity,
      trip
    }})
  }

  return (
    <Scroll>
    <div className="relative w-full h-screen">
      <HeaderLogo />

      <div className="absolute z-50 flex flex-col gap-8 px-6 py-4">
        <FieldBuddy setState={wrapperSetActivity} />
        {isBuddy && (
          <div id="form-buddy" className="">
            <FieldDescription setState={wrapperSetDescription}/>
            <ButtonAction text="Create" onClick={nextCreate} />
          </div>
        )}
      </div>

      <div className="bg-buddy2 h-[460px] w-96 bg-contain bg-center bg-no-repeat absolute bottom-0"></div>
    </div>
    </Scroll>
  )
}

export default Buddy