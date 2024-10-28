import { useCallback, useState } from "react"
import ButtonAction from "../components/buttons/ButtonAction"
import FieldActivity from "../components/forms/FieldActivity"
import FieldDate from "../components/forms/FieldDate"
import FieldDescription from "../components/forms/FieldDescription"
import FieldStyle from "../components/forms/FieldStyle"
import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"


function Preference3() {
  const [activity, setActivity] = useState("")
  const [description, setDescription] = useState(null)
  const [trip, setTrip] = useState("")
  
  const wrapperSetActivity = useCallback(val => {
    setActivity(val);
  }, [setActivity]);
  const wrapperSetDescription = useCallback(val => {
    setDescription(val);
  }, [setDescription]);
  const wrapperSetTrip = useCallback(val => {
    setTrip(val);
  }, [setTrip]);
  
  return (
    <Scroll>
      <HeaderLogo />

      <div className="flex flex-col gap-8 px-6 py-4 text-sm bg-white h-fit w-96">
        <FieldActivity parentStateSetter={wrapperSetActivity} />
          
        <FieldDate />

        <FieldDescription setState={wrapperSetDescription} />
      
        <FieldStyle parentStateSetter={wrapperSetTrip} />
      
        <div className="self-center">
          <ButtonAction text="Confirm" />
        </div>

      </div>
    </Scroll>
  )
}

export default Preference3