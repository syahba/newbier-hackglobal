import { useCallback, useState } from "react"
import ButtonAction from "../components/buttons/ButtonAction"
import FieldActivity from "../components/forms/FieldActivity"
import FieldDate from "../components/forms/FieldDate"
import FieldDescription from "../components/forms/FieldDescription"
import FieldStyle from "../components/forms/FieldStyle"
import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"
import { useNavigate } from "react-router-dom"


function Preference3() {
  const [activity, setActivity] = useState("")
  const [description, setDescription] = useState(null)
  const [trip, setTrip] = useState("")

  const navigate = useNavigate()
  
  const wrapperSetActivity = useCallback(val => {
    setActivity(val);
  }, [setActivity]);
  const wrapperSetDescription = useCallback(val => {
    setDescription(val);
  }, [setDescription]);
  const wrapperSetTrip = useCallback(val => {
    setTrip(val);
  }, [setTrip]);

  const action = async () => {
    let finder
    try {
      var userId = localStorage.getItem("user")
      finder = await fetch("http://localhost:8000/api/itinerary/buddy/finder", {method: "POST", body: JSON.stringify({
        description: description,
        activity: activity,
        trip: trip,
        created_by: parseInt(userId)
      }),
      headers: {
        'Content-Type': 'application/json' // Tipe konten JSON
      }});
      finder = await finder.json()
    } catch (error) {
      console.log(error)
    } finally {
      navigate("/match", {state: {
        activity: activity,
        trip: trip,
        finderId: finder.ID
      }})
    }
  }
  
  return (
    <Scroll>
      <HeaderLogo />

      <div className="flex flex-col gap-8 px-6 py-4 text-sm bg-white h-fit w-96">
        <FieldActivity parentStateSetter={wrapperSetActivity} />
          
        <FieldDate />

        <FieldDescription setState={wrapperSetDescription} />
      
        <FieldStyle parentStateSetter={wrapperSetTrip} />
      
        <div className="self-center">
          <ButtonAction text="Confirm" onClick={action} />
        </div>

      </div>
    </Scroll>
  )
}

export default Preference3