function FieldActivity({parentStateSetter}) {
  const handleActivityChange = (event) => {
    parentStateSetter(event.target.value);
  };
  
  return (
    <div>
      <h5 className="mb-4 text-sm font-bold">
        Let’s start with your preferred activities!
      </h5>
      
      <div className="flex flex-wrap gap-2">
        <div>
          <input
            type="radio"
            name="activity"
            value="sightseeing"
            id="activity-sightseeing"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-sightseeing"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Sightseeing
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            value="culinary_experiences"
            id="activity-culinary-experiences"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-culinary-experiences"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Culinary Experiences
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            value="outdoor_activities"
            id="activity-outdoor-activities"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-outdoor-activities"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Outdoor Activities
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            value="adventure"
            id="activity-adventure"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-adventure"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Adventure
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            value="cultural_experiences"
            id="activity-cultural-experiences"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-cultural-experiences"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Cultural Experiences
          </label>
        </div>
        <div>
          <input
            type="radio"
            name="activity"
            value="shopping"
            id="activity-shopping"
            className="hidden peer"
            onChange={handleActivityChange}
          />
          <label
            htmlFor="activity-shopping"
            className="inline-flex px-4 py-3 border rounded-lg border-blue peer-checked:bg-blue peer-checked:text-white"
          >
            Shopping
          </label>
        </div>
      </div>
    </div>
  );
}

export default FieldActivity;
