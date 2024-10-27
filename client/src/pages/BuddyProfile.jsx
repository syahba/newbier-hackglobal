import { useCallback, useState } from "react"
import FieldBuddy from "../components/forms/FieldBuddy"
import FieldDescription from "../components/forms/FieldDescription"
import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"
import CardProfile from "../components/forms/CardProfile"
import ButtonAction from "../components/buttons/ButtonAction"
import { useLocation, useNavigate } from "react-router-dom"

function BuddyProfile() {
  const { state } = useLocation();
  const { itinerary, isBuddy, description } = state
  const navigate = useNavigate()

  
  const buttonConfirm = () => {
    navigate("/bridging/3")
  }
  return (
    <Scroll>
    <div className="relative w-full h-screen">
      <HeaderLogo />

      <div className="absolute z-50 flex flex-col items-center gap-8 px-6 py-4">
        <FieldBuddy disabled={true} value={isBuddy && "yes"} />

      <div>
        <h5 className="mb-1 text-sm font-bold">Tell us a bit about yourself!</h5>
        <p className="mb-4 text-xs">
          Writing about your planned activities, such as event you’d like to
          attend or rides you’d like to try, would help us get you a better match.
        </p>

        <CardProfile description={description} />
      </div>

        <ButtonAction text="Confirm" onClick={buttonConfirm} />
      </div>

      <div className="bg-buddy2 h-[460px] w-96 bg-contain bg-center bg-no-repeat absolute bottom-0"></div>
    </div>
    </Scroll>
  )
}

export default BuddyProfile