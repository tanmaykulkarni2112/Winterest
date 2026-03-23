import Button from "./SimButton";
import SignupPage from "../pages/Signup";
import LoginPage from "../pages/Login";
import { useState } from "react";
import SearchBar from "./SearchBar";


const Appbar = () => {
    const [openSignup,SetOpenSignup] = useState(false);
    const [openLogIn,setOpenLogIn]=useState(false);
    const [loggedIn,setLoggedIn]=useState(false);

  return (
    <>  
        {openSignup && <SignupPage 
        onClose={()=>SetOpenSignup(false)} 
        openLogin={()=>{
            SetOpenSignup(false);
            setOpenLogIn(true);
        }}/>}

        {openLogIn && <LoginPage 
        onClose={()=> setOpenLogIn(false)}
        openSignup={()=>{
            setOpenLogIn(false);
            SetOpenSignup(true);
        }}
        />}

        <div className="h-20 w-full flex justify-between border-gray-200 shadow">
            <div className="flex justify-around p-6 px-11 gap-x-2">
                <div className="text-lg text-red-500 hover:text-red-600 font-bold p-2 ">
                    <i>Wintrest</i>
                </div>
                <div className="text-lg font-bold p-2">
                    <div onClick={()=>{
                        // testing purposes
                        setLoggedIn(prev=>!prev)
                    }} className=" hover:bg-slate-100 rounded-lg">
                        Explore
                    </div>
                </div>
            </div>
            <div>
                {loggedIn?(<div className="flex items-center p-6">
                    <SearchBar/>
                </div>):
                <div></div>
                }
            </div>
            <div>
                <div className="flex justify-around p-6 px-11">
                    {/* search Barr or buttons*/}
                    {loggedIn?(<div>
                    </div>):
                    <div className="flex justify-around">
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            About 
                        </div>
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            Businesses
                        </div>
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            Create
                        </div>
                        <div className="text-lg font-bold p-2 hover:bg-slate-100 rounded-lg">
                            News 
                        </div>
                    </div>
                    }
                    

                    {/* buttons or avatar */}
                    <div className="flex gap-x-3 ml-1 me-1">{loggedIn?(
                            <div className="flex gap-x-3 items-center">
                            <div className="w-10 h-10 rounded-full flex justify-center items-center cursor-pointer bg-gray-300">
                                👤
                            </div>
                            <div className="w-10 h-10 rounded-lg flex justify-center items-center cursor-pointer">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="size-6">
                                <path fillRule="evenodd" d="M11.828 2.25c-.916 0-1.699.663-1.85 1.567l-.091.549a.798.798 0 0 1-.517.608 7.45 7.45 0 0 0-.478.198.798.798 0 0 1-.796-.064l-.453-.324a1.875 1.875 0 0 0-2.416.2l-.243.243a1.875 1.875 0 0 0-.2 2.416l.324.453a.798.798 0 0 1 .064.796 7.448 7.448 0 0 0-.198.478.798.798 0 0 1-.608.517l-.55.092a1.875 1.875 0 0 0-1.566 1.849v.344c0 .916.663 1.699 1.567 1.85l.549.091c.281.047.508.25.608.517.06.162.127.321.198.478a.798.798 0 0 1-.064.796l-.324.453a1.875 1.875 0 0 0 .2 2.416l.243.243c.648.648 1.67.733 2.416.2l.453-.324a.798.798 0 0 1 .796-.064c.157.071.316.137.478.198.267.1.47.327.517.608l.092.55c.15.903.932 1.566 1.849 1.566h.344c.916 0 1.699-.663 1.85-1.567l.091-.549a.798.798 0 0 1 .517-.608 7.52 7.52 0 0 0 .478-.198.798.798 0 0 1 .796.064l.453.324a1.875 1.875 0 0 0 2.416-.2l.243-.243c.648-.648.733-1.67.2-2.416l-.324-.453a.798.798 0 0 1-.064-.796c.071-.157.137-.316.198-.478.1-.267.327-.47.608-.517l.55-.091a1.875 1.875 0 0 0 1.566-1.85v-.344c0-.916-.663-1.699-1.567-1.85l-.549-.091a.798.798 0 0 1-.608-.517 7.507 7.507 0 0 0-.198-.478.798.798 0 0 1 .064-.796l.324-.453a1.875 1.875 0 0 0-.2-2.416l-.243-.243a1.875 1.875 0 0 0-2.416-.2l-.453.324a.798.798 0 0 1-.796.064 7.462 7.462 0 0 0-.478-.198.798.798 0 0 1-.517-.608l-.091-.55a1.875 1.875 0 0 0-1.85-1.566h-.344ZM12 15.75a3.75 3.75 0 1 0 0-7.5 3.75 3.75 0 0 0 0 7.5Z" clipRule="evenodd" />
                                </svg>
                            </div>
                            </div>
                        ):(
                            <>
                        <div>
                            <Button onClick={()=>{
                                setOpenLogIn(true);
                            }} text="Log In " colour="bg-red-600" textColour="text-white"></Button>
                        </div>
                        <div>
                            <Button onClick={()=>{
                                SetOpenSignup(true)
                            }} text="Sign Up" colour="bg-slate-200" textColour="text-black"></Button>
                        </div>
                        </>
                        )
                        }
                    </div>
                </div>
            </div>
        </div>
    </>
  );
};

export default Appbar;  