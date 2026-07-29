import MainPage from "./pages/MainPage"
import Footer from "./pages/Footer";
import {BrowserRouter, Routes, Route} from 'react-router'
import ProductPage from "./pages/ProductPage";
import {AuthProvider} from "./context/AuthContext";
import ProfilePage from "./pages/ProfilePage";
import Cart from "./pages/Cart";
import AuthPage from "./pages/AuthPage";
import RegistrationPage from "./pages/RegistrationPage";
import Header from "./pages/Header";
import ConstructorPage from "./pages/ConstructorPage";
import CheckoutPage from "./pages/CheckoutPage";
import CheckoutComplete from "./pages/CheckoutComplete";
import AdminPage from "./pages/AdminPage";
import PageNotFound from "./features/404NotFound/PageNotFound";
import LoadingPage from "./features/LoadingScreen/LoadingScreen";
import {useEffect, useState} from "react";



const App = () => {

    const [appReady, setAppReady] = useState(false);

    useEffect(() => {
        setTimeout(() => setAppReady(true), 2000);
    }, []);

    if (!appReady) {
        return <LoadingPage/>;
    }


        return (
        <AuthProvider>
        <BrowserRouter>
            <div className="App">
            <Header/>
        <Routes>
            <Route path = "/auth" element = {<AuthPage/>} />
            <Route path = "/" element = {<MainPage/>} />
            <Route path = "/auth/registry" element ={<RegistrationPage/>} />
            <Route path = "/item/:id" element = {<ProductPage/>} />
            <Route path ="/profile/" element ={<ProfilePage/>} />
            <Route path = "/cart/" element = {<Cart/>} />
            <Route path = "/constructor/:id" element ={<ConstructorPage/>} />
            <Route path = "/checkout" element ={<CheckoutPage/>} />
            <Route path ="/checkout/applications/:id" element = {<CheckoutComplete/>} />
            <Route path ="/admin" element = {<AdminPage/>} />
            <Route path ="*" element = {<PageNotFound />} />
        </Routes>
            <Footer />
            </div>
        </BrowserRouter>
        </AuthProvider>
    )
}

export default App