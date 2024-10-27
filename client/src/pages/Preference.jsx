import ButtonAction from "../components/buttons/ButtonAction"
import FieldDate from "../components/forms/FieldDate"
import FieldStyle from "../components/forms/FieldStyle"
import HeaderLogo from "../components/headers/HeaderLogo"
import Main from "../layouts/Main"

function Preference() {


  
  return (
    <Main>
      
    <div className="bg-preference1 h-3/4 w-full bg-contain bg-center bg-no-repeat">
      <HeaderLogo />
      <div className="mb-4 mt-12 px-6">
        <h5 className="mb-4 text-sm font-bold">Hi, <span className="text-blue">User</span>! Where do you want to go today?</h5>
        <input type="text" className="mb-4 h-8 w-full px-3 shadow-md rounded-lg" placeholder="Enter a destination" />
      </div>
    </div>

    <div>
      <div className="absolute bottom-0 flex h-1/2 w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
        <FieldDate />
      
        <FieldStyle />
      
        <div className="self-center">
          <ButtonAction text="Confirm" />
        </div>
      </div>
    </div>
  </Main>
  )
}

export default Preference