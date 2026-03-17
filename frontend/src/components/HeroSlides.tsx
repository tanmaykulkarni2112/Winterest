import { useEffect, useState } from "react"
import { AnimatePresence , motion } from "framer-motion";
import Button from "./SimButton";
import food1 from "../assets/food1.jpg";
import food2 from "../assets/food2.png";
import decor3 from "../assets/decor3.jpg";
import decor2 from "../assets/decor2.png";
import fashion1 from "../assets/fashion1.png";
import fashion2 from "../assets/fashion2.png";

import LoginPage from "../pages/Login";
import SignupPage from "../pages/Signup";

const Slides = [
  {
    text: "Weeknight Dinner",
    color: "text-purple-500",
    images: [food1,food2]
  },
  {
    text: "Home Decor",
    color: "text-green-500",
    images: [decor3,decor2]
  },
  {
    text: "Outfit Ideas",
    color: "text-orange-500",
    images: [fashion1,fashion2]
  }
];

const HeroSlides = ()=>{

    const[loginOpen,setLoginOpen]=useState(false);
    const[SigninOpen,setSignInOpen]=useState(false);

    const [index,SetIndex]=useState<number>(0);
    
    useEffect(()=>{
        
        const id:any = setInterval(()=>{
            SetIndex((prev)=>(prev+1 )% Slides.length)
        },4000)

        return ()=>{ clearInterval(id)};
    },[])

    const slide = Slides[index];

    return(
        <section className="flex items-center justify-center gap-30 px-20 h-screen"> 
            {SigninOpen && <SignupPage onClose={()=>{
                setSignInOpen(false)
            }}/>}
            {loginOpen && <LoginPage onClose={()=>{
                setLoginOpen(false)
            }}/>}
            <div className="max-w-xl">
                <h1 className="text-6xl font-bold">
                    Find ideas for
                </h1>
                <AnimatePresence mode="wait">
                    <motion.h2
                        key={slide.text}
                        initial={{ opacity: 0, y: 40 }}
                        animate={{ opacity: 1, y: 0 }}
                        exit={{ opacity: 0, y: -40 }}
                        transition={{ duration: 0.5 }}
                        className={`text-6xl font-bold ${slide.color}`}
                        >
                            {slide.text}
                    </motion.h2>
                </AnimatePresence>

                <div className="flex gap-2 mt-6">
                {Slides.map((_, i) => (
                    <div
                    key={i}
                    className={`w-2 h-2 rounded-full ${
                    i === index ? "bg-black" : "bg-gray-300"
                    }`}
                    />
                ))}
                </div>
                <div className="flex gap-x-3 mt-8">
                    <div>
                    <Button text="Join Wintrest for free" colour="bg-red-500" textColour="text-white" onClick={()=>{
                        setSignInOpen(true);
                    }}/>
                    </div>
                    <div onClick={()=>{
                        setLoginOpen(true);
                    }} className="rounded-xl p-2 text-lg mb-2 hover:bg-gray-200 ">
                        I already have an account 
                    </div>
                </div>
            </div> 

            <div className="relative w-105 h-100 ">
                <AnimatePresence mode="wait">
                    <motion.img
                    key={slide.images[0]}
                    src={slide.images[0]}
                    className="absolute w-[320px] rounded-2xl"
                    initial={{ opacity: 0, rotate: -10 }}
                    animate={{ opacity: 1, rotate: 0 }}
                    exit={{ opacity: 0 }}
                    transition={{ duration: 0.7 }}
                    />
                </AnimatePresence>

                <AnimatePresence mode="wait">
                    <motion.img
                        key={slide.images[1]}
                        src={slide.images[1]}
                        className="absolute w-45 -bottom-5 -right-15 rounded-2xl rotate-12 shadow-xl"
                        initial={{ opacity: 0, y: 40 }}
                        animate={{ opacity: 1, y: 0 }}
                        exit={{ opacity: 0 }}
                        transition={{ duration: 0.7 }}
                    />
                    </AnimatePresence>
            </div>
        </section>  
    )
}

export default HeroSlides;