
import {memo} from "react";
import {useParams} from "react-router";
import {GetItemCard} from "../components/ProductPage/hooks/getItemCard";
import ProductInfoHeader from "../components/ProductPage/components/ProductInfoHeader";
import ProductSpecs from "../components/ProductPage/components/ProductSpecs";
import ProductInfo from "../components/ProductPage/components/ProductInfo";
import {useAuth} from "../context/AuthContext";
import PageNotFound from "../features/404NotFound/PageNotFound";



const ProductPage = () => {
    const {id} = useParams()
    const itemCard = GetItemCard(id)
    const user = useAuth()



    if (itemCard.loading) {
        return <div>loading...</div>


    }
    if (itemCard.error) {
        return <PageNotFound />
    }
    return (

        <div className ="product-info-wrapper">
            {(user.userRole === "admin" && !itemCard.item.itemData.visible) || ((user.userRole !=="admin" || user.userRole ==="admin" ) && itemCard.item.itemData.visible) ? (
            <div className="product-info">
                <ProductInfoHeader
                    itemCard = {itemCard.item}
                    userRole = {user.userRole}
                    />
              <ProductInfo
              itemCard = {itemCard.item}
              />
                <ProductSpecs
                    itemCard={itemCard.item}
                    userRole = {user.userRole}
                    setItem = {itemCard.setItem}
                    />
            </div>
                ) : (
                    <h1>Такого предмета не существует</h1>
            )}
        </div>
    )
}

export default memo(ProductPage);