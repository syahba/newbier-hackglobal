import { useEffect, useState } from "react";

function FieldDescription({setState}) {
  const [description, setDescription] = useState()

  useEffect(()=> {
    setState(description)
  }, [description])
  

  return (
    <div>
      <h5 className="mb-1 text-sm font-bold">Tell us a bit about yourself!</h5>
      <p className="mb-4 text-xs">
        Writing about your planned activities, such as event you’d like to
        attend or rides you’d like to try, would help us get you a better match.
      </p>
      <textarea
        name="desc"
        id="desc"
        placeholder="Write your interests"
        value={description}
        onChange={e => setDescription(e.target.value)}
        className="w-full h-32 p-3 text-sm rounded-lg shadow-lg outline-blue"
      ></textarea>
    </div>
  );
}

export default FieldDescription;
