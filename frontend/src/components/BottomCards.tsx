import { useState, type JSX } from "react";
import card1 from "../assets/card1.png";
import card2 from "../assets/card2.png";
import card3 from "../assets/card3.png";
import Button from "./SimButton";

import SignupPage from "../pages/Signup";
import LoginPage from "../pages/Login";

const cards = [
  {
    title: "Collaborate with group boards",
    desc: "Visualize your ideas with others using a Wintrest account.",
    img: card1,
  },
  {
    title: "Save ideas you love",
    desc: "Collect inspiration and organize everything in one place.",
    img: card2,
  },
  {
    title: "Discover new inspiration",
    desc: "Explore ideas shared by people around the world.",
    img: card3,
  },
];

const BottomCards = (): JSX.Element => {
  const [openSignUp,setOpenSignUp]=useState(false);
  const [openLogin,setOpenLogin]=useState(false);
  return (
    <div className="bg-gray-100 w-full py-28">
      
      {openSignUp && <SignupPage onClose={()=>{
        setOpenSignUp(false)
      }}
      openLogin={()=>{
        setOpenSignUp(false);
        setOpenLogin(true);
      }}
      />}

      {openLogin && <LoginPage
      onClose={()=>setOpenLogin(false)}
      openSignup={()=>{
        setOpenLogin(false);
        setOpenSignUp(true);
      }}
      />}

      <div className="max-w-5xl mx-auto flex flex-col items-center text-center gap-6 px-6">
        <div className="text-5xl md:text-6xl font-bold">
          Bring your favorite ideas to life
        </div>

        <div className="text-xl text-gray-600">
          With Wintrest, you can use tools that spark your creativity and help
          you find more inspiration.
        </div>
      </div>

      <div className="max-w-6xl mx-auto mt-28 flex flex-col gap-32 px-6">

        {cards.map((card, index) => (
          <div
            key={index}
            className={`flex items-center gap-16 ${
              index % 2 === 0 ? "flex-row" : "flex-row-reverse"
            }`}
          >

            <div className="w-1/2 flex justify-center">
              <img
                src={card.img}
                className="rounded-2xl shadow-lg w-105 transition duration-300 hover:scale-105 hover:shadow-2xl"
                alt="feature"
              />
            </div>

            <div className="w-1/2 flex flex-col gap-6">

              <div className="text-4xl md:text-5xl font-bold">
                {card.title}
              </div>

              <div className="text-xl text-gray-600">
                {card.desc}
              </div>

              <div className="mt-2">
                <div className="inline-block transition hover:scale-105">
                  <Button
                    text="See an example"
                    colour="bg-red-500 hover:bg-red-600"
                    textColour="text-white"
                    onClick={() => {
                      setOpenSignUp(true);
                    }}
                  />
                </div>
              </div>

            </div>

          </div>
        ))}

      </div>
    </div>
  );
};

export default BottomCards;