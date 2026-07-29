import { memo } from "react";

import GetItems from "../components/MainPage/hooks/GetItems";

import MainPageFilter from "../components/MainPage/components/MainPageFilter";
import MainPageGoods from "../components/MainPage/components/MainPageGoods";
import MainPageHeader from "../components/MainPage/components/MainPageHeader";
import PageNotFound from "../features/404NotFound/PageNotFound";
import LoadingScreen from "../features/LoadingScreen/LoadingScreen";
const MainPage = () => {
    const items = GetItems()

    if (items.loading) {
        return <div>loading...</div>
    }

    if (items.error) {
        return <PageNotFound />
    }



    return (
<div className = "catalog-page-wrapper">
    <div className ="catalog-page">
    <MainPageHeader />
    <MainPageFilter changeActive = {items.setActiveCategory}
            activeCategory = {items.activeCategory}
            changeSecondary = {items.setActiveSecondary}
            activeSecondary = {items.activeSecondary}
            secondaryCategory ={items.secondaryTypes}
            types = {items.categoryMap}
    />
    <MainPageGoods itemsList = {items} />
</div>
</div>
    )
}

export default memo(MainPage)
