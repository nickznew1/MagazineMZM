import {useEffect, useMemo, useState} from "react";
import {useAuth} from "../../../context/AuthContext";


function GetItems () {


        const [items, setItems] = useState([])
        const [loading, setLoading] = useState(false)
        const [error, setError] = useState(null)
        const [activeCategory, setActiveCategory] = useState("all");
        const [activeSecondary, setActiveSecondary] = useState(null)
        const user = useAuth()




        const categoryMap = {
            level: {
                title: "Датчики уровня",
                type: "Датчик уровня",
            },
            pressure: {
                title: "Датчики давления",
                type: "Датчик давления",
            },
            flow: {
                title: "Датчики расхода",
                type: "Датчик расхода",
            }
        };


            const secondaryTypes = useMemo(() => {
                if (!activeCategory || activeCategory === "all") return [];
                const currentType = categoryMap[activeCategory]?.type;

                if (user.userRole === "admin") {
                    const filtered = items?.filter(
                        item => item.item_type === currentType);
                    if (filtered) {
                        return [...new Set(filtered?.map(item => item.item_secondary_type))];
                    } else {
                        return []
                    }

                } else {
                    const filtered = items?.filter(
                        item =>item.visible === true && item.item_type === currentType
                    );
                    if (filtered.length > 0) {
                        return [...new Set(filtered?.map(item => item.item_secondary_type))];
                    } else {
                        return []
                    }
                }


            }, [items, activeCategory, user, categoryMap]);

            const filteredItems = useMemo(() => {
                if (activeCategory === "all" ) {
                    return user.userRole === "admin"
                        ? items
                        : items?.filter(item => item.visible === true);
                }


                const currentType = categoryMap[activeCategory]?.type;
                if (user.userRole === "admin") {
                    let filtered = items?.filter(
                        item => item.item_type === currentType
                    );

                    if (activeSecondary) {
                        filtered = filtered?.filter(
                            item =>   item.item_secondary_type === activeSecondary
                        )

                    }
                    return filtered;
                } else {

                    let filtered = items?.filter(
                        item =>  item.item_type  === currentType && item.visible === true
                    );

                    if (activeSecondary ) {
                        filtered = filtered?.filter(
                            item => item.visible === true && item.item_secondary_type === activeSecondary
                        )

                    }
                    return filtered;
                }

            }, [items, activeCategory, activeSecondary, user, categoryMap])


        useEffect(()=>{
            const getItems = async () => {
                try {
                    setLoading(true)
                    setError(null)
                    const response = await fetch('http://localhost:8080/item/all', { credentials: 'include' });
                    if (response.ok) {
                        const itemsList = await response.json();
                        setLoading(false)
                        setItems(itemsList);
                    } else {
                        setLoading(false)
                        setItems(null);
                    }
                } catch (err) {
                    setError(err)
                    setItems(null);
                } finally {
                    setLoading(false)
                }
            };
            getItems()
        }, []);

        return {items, setItems, loading, error, setActiveCategory,activeCategory,activeSecondary, setActiveSecondary, secondaryTypes, categoryMap, filteredItems}
    }

export default GetItems