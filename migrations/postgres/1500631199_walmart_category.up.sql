CREATE TABLE IF NOT EXISTS walmart_category (
    id VARCHAR(30) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

--FOOD--
--Baking
INSERT INTO walmart_category(id, name) 
VALUES('976759_976780_1044119', 'Flours & Meals');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976780_1044152', 'Oils & Shortening');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976780_1007682', 'Salt, Spices & Seasoning');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976780_1044129', 'Sugar & Sweeteners');
--Beverages
INSERT INTO walmart_category(id, name) 
VALUES('976759_976782_1001680', 'Soft Drinks');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976782_1001659', 'Water');
--Coffee
INSERT INTO walmart_category(id, name)
VALUES('976759_1086446_1229650', 'Instant Coffee');
INSERT INTO walmart_category(id, name)
VALUES('976759_1086446_1229654', 'Ready to Drink Coffee');
--Breakfast & Cereal
INSERT INTO walmart_category(id, name) 
VALUES('976759_976783_1231208', 'Cold Cereal');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976783_1231206', 'Pancake & Waffle Mix');
--Candy & Gum
INSERT INTO walmart_category(id, name) 
VALUES('976759_1096070_1224976', 'Chocolate');
INSERT INTO walmart_category(id, name) 
VALUES('976759_1096070_12249975', 'Hard Candy & Lollipop');
--Fresh Food
INSERT INTO walmart_category(id, name) 
VALUES('976759_1071964_976779', 'Bakery & Bread');
INSERT INTO walmart_category(id, name) 
VALUES('976759_1071964_976788', 'Dairy, Eggs & Cheese');
INSERT INTO walmart_category(id, name) 
VALUES('976759_1071964_976796', 'Meat, Seafood & Poultry');
--Frozen Foods
INSERT INTO walmart_category(id, name) 
VALUES('976759_976791_1001413', 'Frozen Fruit');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976791_1001424', 'Frozen Vegetables');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976791_1001423', 'Ice Cream & Novelties');
--Meal Solutions, Grains & Pasta
INSERT INTO walmart_category(id, name) 
VALUES('976759_976794_976785', 'Canned Goods & Soups');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976794_1001474', 'Grains & Rice');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976794_1007519', 'Pasta & Noodles');
--Snacks, Cookies & Chips
INSERT INTO walmart_category(id, name) 
VALUES('976759_976787_1001390', 'Chips');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976787_1001391', 'Cookies');
INSERT INTO walmart_category(id, name) 
VALUES('976759_976787_1001471', 'Pudding & Gelatin');

--Health
--Medicine Cabinet
INSERT INTO walmart_category(id, name) 
VALUES('976760_976798_1001541', 'Cough, Cold & Flu');
INSERT INTO walmart_category(id, name) 
VALUES('976760_976798_1001546', 'First Aid');
INSERT INTO walmart_category(id, name) 
VALUES('976760_976798_1001540', 'Digestion & Nausea');
--Personal Care
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005862_1071969', 'Bath & Body');
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005862_1001487', 'Deodorants & Antiperspirant');
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005862_1007221', 'Oral Care');
--Vitamins
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005863_1225237', 'Children''s Vitamins');
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005863_1001552', 'Herbals');
INSERT INTO walmart_category(id, name) 
VALUES('976760_1005863_1091528', 'Supplements');

--Household Essentials
--Bathroom
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071965_1001719', 'Hand Soap & Sanitizers');
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071965_1149384', 'Toilet paper');
--Cleaning Supplies
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071966_1072133', 'All-Purpose Cleaners');
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071966_1073104', 'Carpet & Floor Cleaners');
--Kitchen
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071968_1155752', 'Dish Soap');
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071968_1149383', 'Paper Towels & Napkins');
--Laundry Room
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071967_1149392', 'Fabric Softeners');
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1071967_1149379', 'Laundry Detergents');
--Paper & Plastic
INSERT INTO walmart_category(id, name) 
VALUES('1115193_1073264_1149389', 'Facial Tissue');












