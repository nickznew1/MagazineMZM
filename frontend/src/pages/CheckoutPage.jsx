
import {GetCheckout} from "../components/CheckoutPage/hooks/GetCheckout";
import CheckoutHeader from "../components/CheckoutPage/components/CheckoutHeader";
import CheckoutItems from "../components/CheckoutPage/components/CheckoutItems";
import CheckoutFields from "../components/CheckoutPage/components/CheckoutFields";
import NotAuthorized from "../features/NotAuthorized/NotAuthorized";



const CheckoutPage = () => {
    const checkoutInfo = GetCheckout();

    if (checkoutInfo.loading){
        return <div>loading...</div>
    }

    if (checkoutInfo.error){
        return <NotAuthorized/>
    }




    return(
        <div className = "checkout-page-wrapper">
            <div className = "checkout-page">
                <CheckoutHeader />
                <CheckoutItems checkout= {checkoutInfo.inputs} />
                <CheckoutFields checkout={checkoutInfo} />
    </div>
        </div>
    )
}

export default CheckoutPage;