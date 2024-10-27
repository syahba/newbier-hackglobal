import { useCallback, useState } from "react"
import FieldBuddy from "../components/forms/FieldBuddy"
import FieldDescription from "../components/forms/FieldDescription"
import HeaderLogo from "../components/headers/HeaderLogo"
import Scroll from "../layouts/Scroll"

function Buddy() {
  const [isBuddy, setIsBuddy] = useState(false)

  const wrapperSetActivity = useCallback(val => {
    setIsBuddy(val);
  }, [setIsBuddy]);

  return (
    <Scroll>
    <div class="relative w-full h-screen">
      <HeaderLogo />

      <div class="absolute z-50 flex flex-col gap-8 px-6 py-4">
        <FieldBuddy setState={wrapperSetActivity} />
        {isBuddy && (
          <div id="form-buddy" class="">
            <FieldDescription />
            <button class="self-center px-5 py-3 mt-2 text-sm rounded-lg shadow-lg w-fit bg-yellow hover:bg-blue">Create</button>
          </div>
        )}
      </div>

      <div class="bg-buddy2 h-[460px] w-96 bg-contain bg-center bg-no-repeat absolute bottom-0"></div>
    </div>
    </Scroll>
  )
}

export default Buddy