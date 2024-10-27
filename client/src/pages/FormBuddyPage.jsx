import ButtonAction from "../components/buttons/ButtonAction";
import FieldBuddy from "../components/forms/FieldBuddy";
import FieldDescription from "../components/forms/FieldDescription";
import HeaderLogo from "../components/headers/HeaderLogo";

function FormBuddyPage() {
  return (
    <div class="w-full h-screen relative">
      <HeaderLogo></HeaderLogo>

      <div class="flex flex-col gap-8 px-6 py-4 z-50 absolute">
        <FieldBuddy></FieldBuddy>

        <div id="form-buddy" class="">
          <FieldDescription></FieldDescription>

          <ButtonAction text={"Create"}></ButtonAction>
        </div>
      </div>

      <div class="bg-buddy2 h-[460px] w-96 bg-contain bg-center bg-no-repeat absolute bottom-0"></div>
    </div>
  );
}

export default FormBuddyPage;
