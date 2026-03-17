import type { JSX } from 'react'
import Appbar from '../components/Appbar'
import BottomCards from '../components/BottomCards'
import BottomInfo from '../components/BottomInfo'
import HeroSlides from '../components/HeroSlides'

const Landing = ():JSX.Element => {
    return(<>
        <Appbar/>
        <HeroSlides/>
        <BottomCards/>
        <BottomInfo/>
    </>);
}

export default Landing;