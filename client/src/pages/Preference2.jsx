import { useState } from "react"
import ButtonAction from "../components/buttons/ButtonAction"
import FieldActivity from "../components/forms/FieldActivity"
import FieldDate from "../components/forms/FieldDate"
import FieldStyle from "../components/forms/FieldStyle"
import HeaderLogo from "../components/headers/HeaderLogo"
import Main from "../layouts/Main"
import React, { useEffect, useCallback, useRef } from "react";
import { useNavigate } from "react-router-dom"

function Preference2() {
  const [activity, setActivity] = useState("")
  const [trip, setTrip] = useState("")
  const navigate = useNavigate()

  const wrapperSetActivity = useCallback(val => {
    setActivity(val);
  }, [setActivity]);
  const wrapperSetTrip = useCallback(val => {
    setTrip(val);
  }, [setTrip]);
  
  const action = () => {
    const fetchData = async () => {
      try {
        console.log(activity, trip)
        const response = await fetch(`${process.env.HOST}/api/generate-itinerary?activity=${activity}&trip=${trip}`);
        const data = await response.json();

        navigate("/itinerary",{state: {
          data: data,
          activity: activity,
          trip: trip,
        }})
      } catch(err) {
        console.log(err)
      }
    }
    fetchData()
  } 

  return (
    <Main>
      <div className="w-full bg-center bg-no-repeat bg-contain h-2/5 bg-preference2">
        <HeaderLogo/>
      </div>

      <div>
        <div className="absolute bottom-0 flex h-3/4 w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
          <FieldActivity parentStateSetter={wrapperSetActivity} />
          
          <FieldDate />
        
          <FieldStyle parentStateSetter={wrapperSetTrip} />
        
          <div className="self-center">
            <ButtonAction text="Confirm" onClick={() => action()} />
          </div>
        </div>
      </div>
    </Main>
  )
}

export default Preference2