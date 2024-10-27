import FieldBuddy from "../components/forms/FieldBuddy";
import FieldDescription from "../components/forms/FieldDescription";
import HeaderLogo from "../components/headers/HeaderLogo";
import ButtonAction from "../components/buttons/ButtonAction";

function ProfileBuddyPage() {
  return (
    <div class="w-full h-screen relative">
      <HeaderLogo></HeaderLogo>

      <div class="flex flex-col gap-8 px-6 py-4 z-50 items-center absolute">
        <FieldBuddy></FieldBuddy>

        <FieldDescription></FieldDescription>

        <ButtonAction text={"Confirm"}></ButtonAction>
      </div>

      <div class="bg-buddy2 h-[460px] w-96 bg-contain bg-center bg-no-repeat absolute bottom-0"></div>
    </div>
  );
}

export default ProfileBuddyPage;
