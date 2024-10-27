import ButtonAction from "../components/buttons/ButtonAction"
import FieldActivity from "../components/forms/FieldActivity"
import FieldDate from "../components/forms/FieldDate"
import FieldStyle from "../components/forms/FieldStyle"
import HeaderLogo from "../components/headers/HeaderLogo"
import Main from "../layouts/Main"

function Preference2() {
  

  return (
    <Main>
      <div class=" w-full h-2/5 bg-preference2 bg-contain bg-center bg-no-repeat">
        <HeaderLogo />
      </div>

      <div>
        <div class="absolute bottom-0 flex h-3/4 w-96 flex-col gap-8 rounded-t-3xl bg-white px-6 pt-5 text-sm shadow-[0_-4px_20px_rgba(0,0,0,0.1)]">
          <FieldActivity />
          
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

export default Preference2