function CardProfile(user) {
  return (
    <div className="flex h-auto w-80 flex-col gap-3 rounded-lg bg-white p-4 text-sm shadow-lg mb-10 mt-5">
      <h5 className="font-bold">Come join my trip to Universal Studio!</h5>
      <div className="flex gap-12">
        <div className="flex gap-2 items-center">
          <div className="w-5 h-5 bg-black"></div>
          <p>{user.name}</p>
        </div>
        <div className="flex gap-2 items-center">
          <div className="w-5 h-5 bg-black"></div>
          <p>{user.country}</p>
        </div>
        <div className="flex gap-2 items-center">
          <div className="w-5 h-5 bg-black"></div>
          <p>22</p>
        </div>
      </div>
      <p>
        I enjoy extreme rides and horror house! I’m planning to try every single
        rides at USS and be spooked at the legendary horror house~
      </p>
    </div>
  );
}

export default CardProfile;
