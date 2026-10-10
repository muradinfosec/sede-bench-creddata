USE cycling;
CREATE KEYSPACE IF NOT EXISTS cycling WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

// Q1:
-- Find a cyclist's name given an ID number
// CREATE TABLE SIMPLE PRIMARY KEY
CREATE TABLE cycling.cyclist_name ( id UUID PRIMARY KEY, lastname text, firstname text );
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (6e4424fd-8c26-8b46-5f97-eaebb2d036f5, 'VOS','Marianne');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (c5ce8089-ac5e-9612-d95c-3008aea3fcfa, 'VAN DER BREGGEN','Anna');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (d0da5cf1-f950-2c56-e204-29865bcd4fe6, 'FRAME','Alex');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (751758ab-8406-52c2-0a0f-0c1c5b75edda, 'TIRALONGO','Paolo');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (5cf65aad-a98a-85e8-c5d6-28d1cd7f6e11, 'KRUIKSWIJK','Steven');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (dc409614-fe05-3ef4-0168-6fd25d743dca, 'MATTHEWS', 'Michael');
SELECT * FROM cycling.cyclist_name;
SELECT lastname, firstname FROM cycling.cyclist_name WHERE id = 1fb76cab-a21e-42a8-f3c5-81a7fa7a6a22;

-- Q2:
-- Find cyclists that fit a particular category
// CREATE TABLE CLUSTERING ORDER, PRIMARY KEY: PARTITION KEY + 1 CLUSTERING COLUMN, SIMPLE WHERE QUERY
CREATE TABLE cycling.cyclist_category ( category text, points int, id UUID, lastname text, PRIMARY KEY (category, points)) WITH CLUSTERING ORDER BY (points DESC);
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('GC',1269,344519bb-3963-48c3-6f7f-0f3b3b05bbdd,'TIRALONGO');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('One-day-races',367,666381ec-1461-07a0-1d3d-8f4d8d24deff,'TIRALONGO');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('Time-trial',182,465836ad-4856-71c4-3c8d-0c7b6b45ccbf,'TIRALONGO');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('Sprint',0,740002dc-3309-14c3-0f0c-9c9a0a07cafe,'TIRALONGO');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('GC',1324,1da09ecc-c99c-70e4-b3f2-46e3cf7a6f60,'KRUIJSWIJK');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('One-day-races',198,1fc30dfc-b20d-82b1-d2f6-35d0ae3b9a81,'KRUIJSWIJK');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('Sprint',39,1ad21bdc-c33b-13e8-a7b0-96d5ec5e6c05,'KRUIJSWIJK');
INSERT INTO cycling.cyclist_category (category, points, id, lastname) VALUES ('Time-trial',3,9fb30fab-f61b-15c3-a6f1-16c7de5c3c13,'KRUIJSWIJK');
SELECT * FROM cycling.cyclist_category;
SELECT lastname, points FROM cycling.cyclist_category WHERE category = 'One-day-races';

-- Q3:
-- Store race information by year and race name using a COMPOSITE PARTITION KEY
CREATE TABLE cycling.rank_by_year_and_name ( race_year int, race_name text, cyclist_name text, rank int, PRIMARY KEY ((race_year, race_name), rank) );
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Benjamin PRADES', 1);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Adam PHELAN', 2);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Thomas LEBAS', 3);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2015, 'Giro d''Italia - Stage 11 - Forli > Imola', 'Ilnur ZAKARIN', 1);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2015, 'Giro d''Italia - Stage 11 - Forli > Imola', 'Carlos BETANCUR', 2);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2014, '4th Tour of Beijing', 'Phillippe GILBERT', 1);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2014, '4th Tour of Beijing', 'Daniel MARTIN', 2);
INSERT INTO cycling.rank_by_year_and_name (race_year, race_name, cyclist_name, rank) VALUES (2014, '4th Tour of Beijing', 'Johan Esteban CHAVES', 3);
SELECT * FROM cycling.rank_by_year_and_name;
SELECT * FROM cycling.rank_by_year_and_name WHERE race_year=2015 AND race_name='Tour of Japan - Stage 4 - Minami > Shinshu';

-- New C* 3.6
-- PER PARTITION LIMIT
-- To get the Top Two for each race_year-race_name pair
SELECT * FROM cycling.rank_by_year_and_name PER PARTITION LIMIT 2;

-- Q4:
-- Find a cyclist's id given lastname and firstname
-- Another CREATE TABLE using COMPOSITE PARTITION KEY
-- 2i INDEX ALSO GOOD FOR THIS TABLE
CREATE TABLE cycling.cyclist_id ( lastname text, firstname text, age int, id UUID, PRIMARY KEY ((lastname, firstname), age) );
INSERT INTO cycling.cyclist_id (lastname, firstname, age, id) VALUES ('EENKHOORN','Pascal',18, cbfad8d4-9ed3-53b3-cfbd-5feadabd6910);
INSERT INTO cycling.cyclist_id (lastname, firstname, age, id) VALUES ('WELTEN','Bram',18, 11a027ec-c424-3df3-e5b4-a9d9be2ba069);
INSERT INTO cycling.cyclist_id (lastname, firstname, age, id) VALUES ('COSTA','Adrien',17, 74b179ff-e785-3ad6-fe1d-2a5648916541);
SELECT * FROM cycling.cyclist_id WHERE lastname = 'COSTA' AND firstname = 'Adrien';
-- If you want to search by age, an index can be added
CREATE INDEX c_age ON cycling.cyclist_id (age);
SELECT * FROM cycling.cyclist_id WHERE age = 18;

-- Q5:
-- Display flag for riders
-- CREATE TABLE WITH STATIC COLUMN, example uses an integer to identify flag, but it could be a blob
CREATE TABLE cycling.country_flag (country text, cyclist_name text, flag int STATIC, PRIMARY KEY (country, cyclist_name));
INSERT INTO cycling.country_flag (country, cyclist_name, flag) VALUES ('Belgium', 'Jacques', 1);
INSERT INTO cycling.country_flag (country, cyclist_name) VALUES ('Belgium', 'Andre');
INSERT INTO cycling.country_flag (country, cyclist_name, flag) VALUES ('France', 'Andre', 2);
INSERT INTO cycling.country_flag (country, cyclist_name, flag) VALUES ('France', 'George', 3);
-- USE SELECT REPEATEDLY TO SHOW CHANGING (OR UNCHANGING) NATURE OF the column 'flag'
SELECT * FROM cycling.country_flag;

-- Q6:
-- Find all teams that a cyclist has been a member of
--CREATE TABLE WITH SET
CREATE TABLE cycling.cyclist_career_teams ( id UUID PRIMARY KEY, lastname text, teams set<text> );
INSERT INTO cycling.cyclist_career_teams (id,lastname,teams) VALUES (9e1112bb-4d07-7c49-3d14-babec0b921f8, 'VOS', { 'Rabobank-Liv Woman Cycling Team','Rabobank-Liv Giant','Rabobank Women Team','Nederland bloeit' } );
INSERT INTO cycling.cyclist_career_teams (id,lastname,teams) VALUES (f2ac1780-cf9e-4430-c28b-2702dfa8dfad, 'VAN DER BREGGEN', { 'Rabobank-Liv Woman Cycling Team','Sengers Ladies Cycling Team','Team Flexpoint' } );
INSERT INTO cycling.cyclist_career_teams (id,lastname,teams) VALUES (ae42baaa-abd7-6d26-d19d-cfdc20e0de54, 'ARMITSTEAD', { 'Boels-Dolmans Cycling Team','AA Drink - Leontien.nl','Team Garmin - Cervelo' } );
INSERT INTO cycling.cyclist_career_teams (id,lastname,teams) VALUES (8e4ddc70-6bee-7ba9-df83-bbb231983a58, 'BRAND', { 'Rabobank-Liv Woman Cycling Team','Rabobank-Liv Giant','AA Drink - Leontien.nl','Leontien.nl' } );
SELECT lastname,teams FROM cycling.cyclist_career_teams;
SELECT lastname, teams FROM cycling.cyclist_career_teams WHERE id=0c0746df-3e57-2d55-6b19-ecebd6c013d2;

-- NOT A QUERY, JUST A TABLE FOR QUERIES
-- CREATE TABLE WITH LIST FOR UPDATE
-- The SELECT statements that use this table can be found below
CREATE TABLE cycling.calendar (race_id int, race_name text, race_start_date timestamp, race_end_date timestamp, PRIMARY KEY (race_id, race_start_date, race_end_date));
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (100, 'Giro d''Italia','2015-05-09','2015-05-31');
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (101, 'Criterium du Dauphine','2015-06-07','2015-06-14');
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (102, 'Tour de Suisse','2015-06-13','2015-06-21');
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (103, 'Tour de France','2015-07-04','2015-07-26');
SELECT * FROM cycling.calendar;

-- NEW FOR C*3.6
-- Clustering columns can be used in a WHERE clause with ALLOW FILTERING without secondary indexes
-- This query uses the clustering column "race_start_date" without an index and without using the partition key
-- but using ALLOW FILTERING
SELECT * FROM cycling.calendar WHERE race_start_date='2015-06-13' ALLOW FILTERING;

-- Q7:
-- Find all calendar events for a particular year and month
CREATE TABLE cycling.upcoming_calendar ( year int, month int, events list<text>, PRIMARY KEY ( year, month ));
INSERT INTO cycling.upcoming_calendar (year, month, events) VALUES (2015, 06, ['Criterium du Dauphine','Tour de Suisse']);
INSERT INTO cycling.upcoming_calendar (year, month, events) VALUES (2015, 07, ['Tour de France']);
SELECT * FROM cycling.upcoming_calendar WHERE year=2015 AND month=06;

-- Q8:
-- SIMPLE USER-DEFINED TYPE
CREATE TYPE cycling.fullname ( firstname text, lastname text );
CREATE TABLE cycling.race_winners (race_name text, race_position int, cyclist_name FROZEN<fullname>, PRIMARY KEY (race_name, race_position));
INSERT INTO cycling.race_winners (race_name, race_position, cyclist_name) VALUES ('National Championships South Africa WJ-ITT (CN)', 1, {firstname:'Frances',lastname:'DU TOUT'});
INSERT INTO cycling.race_winners (race_name, race_position, cyclist_name) VALUES ('National Championships South Africa WJ-ITT (CN)', 2, {firstname:'Lynette',lastname:'BENSON'});
INSERT INTO cycling.race_winners (race_name, race_position, cyclist_name) VALUES ('National Championships South Africa WJ-ITT (CN)', 3, {firstname:'Anja',lastname:'GERBER'});
INSERT INTO cycling.race_winners (race_name, race_position, cyclist_name) VALUES ('National Championships South Africa WJ-ITT (CN)', 4, {firstname:'Ame',lastname:'VENTER'});
INSERT INTO cycling.race_winners (race_name, race_position, cyclist_name) VALUES ('National Championships South Africa WJ-ITT (CN)', 5, {firstname:'Danielle',lastname:'VAN NIEKERK'});
SELECT * FROM cycling.race_winners WHERE race_name = 'National Championships South Africa WJ-ITT (CN)';

-- Q9:
-- Find all races for a particular cyclist
-- CREATE TYPE - User-Defined Type, race
-- CREATE TABLE WITH LIST, SIMPLE PRIMARY KEY
CREATE TYPE cycling.race (race_title text, race_date timestamp, race_time text);
CREATE TABLE cycling.cyclist_races ( id UUID PRIMARY KEY, lastname text, firstname text, races list<FROZEN <race>> );
INSERT INTO cycling.cyclist_races (id, lastname, firstname, races) VALUES (5c3631ec-4c91-9c69-3c42-acfdb9a402c7, 'VOS', 'Marianne', [ {race_title:'Rabobank 7-Dorpenomloop Aalburg',race_date:'2015-05-09',race_time:'02:58:33'},{race_title:'Ronde van Gelderland',race_date:'2015-04-19',race_time:'03:22:23'}
]);
INSERT INTO cycling.cyclist_races (id, lastname, firstname, races) VALUES (e1ff1972-fd2b-3042-a64b-5059def9bdad, 'VAN DER BREGGEN', 'Anna', [ {race_title:'Festival Luxembourgeois du cyclisme feminin Elsy Jacobs - Prologue - Garnich > Garnich',race_date:'2015-05-01',race_time:'08:13:00'},{race_title:'Fest
ival Luxembourgeois du cyclisme feminin Elsy Jacobs - Stage 2 - Garnich > Garnich',race_date:'2015-05-02',race_time:'02:41:52'},{race_title:'Festival Luxembourgeois du cyclisme feminin Elsy Jacobs - Stage 3 - Mamer > Mamer',race_date:'2015-05-03',race_time:'02:31:24'} ]);
SELECT * FROM cycling.cyclist_races;
SELECT lastname, races FROM cycling.cyclist_races WHERE id = e5cb5435-fb5b-5863-f61f-5610cde6afbf;

-- Q10:
-- Find all teams for a particular cyclist associated with the year of membership
-- teams map<int, text> is map<year, team_name>
-- CREATE TABLE WITH MAP, SIMPLE PRIMARY KEY
CREATE TABLE cycling.cyclist_teams ( id UUID PRIMARY KEY, lastname text, firstname text, teams map<int,text> );
INSERT INTO cycling.cyclist_teams (id, lastname, firstname, teams) VALUES (2b9372ac-9c38-8f42-5b75-bcdba2d906f5,'VOS', 'Marianne', {2015 : 'Rabobank-Liv Woman Cycling Team', 2014 : 'Rabobank-Liv Woman Cycling Team', 2013 : 'Rabobank-Liv Giant', 2012 : 'Rabobank Women Team', 2011 : 'Nederland bloeit' });
INSERT INTO cycling.cyclist_teams (id, lastname, firstname, teams) VALUES (a8ca9413-fc9b-9014-c61d-1268ffd6ddaf,'VAN DER BREGGEN', 'Anna', {2015 : 'Rabobank-Liv Woman Cycling Team', 2014 : 'Rabobank-Liv Woman Cycling Team', 2013 : 'Sengers Ladies Cycling Team', 2012 : 'Sengers Ladies Cycling Team', 2009 : 'Team Flexpoint' });
INSERT INTO cycling.cyclist_teams (id, lastname, firstname, teams) VALUES (cb43abfe-dcd1-4f99-e53b-defd23b3fb13,'ARMITSTEAD', 'Elizabeth', {2015 : 'Boels-Dolmans Cycling Team', 2014 : 'Boels-Dolmans Cycling Team', 2013 : 'Boels-Dolmans Cycling Team', 2012 : 'AA Drink - Leontien.nl', 2011 : 'Team Garmin - Cervelo' });
SELECT lastname, firstname, teams FROM cycling.cyclist_teams;
SELECT lastname, firstname, teams FROM cycling.cyclist_teams WHERE id=0f7855ad-3f53-7c83-9c41-cdbbb2f978a2;

-- Q11:
-- Find all stats for a particular cyclist
-- CREATE TYPE -  UDT, basic_info
-- CREATE TABLE with UDT, SIMPLE PRIMARY KEY
CREATE TYPE cycling.basic_info ( birthday timestamp, nationality text, weight text, height text );
CREATE TABLE cycling.cyclist_stats ( id UUID, lastname text, basics FROZEN <basic_info>, PRIMARY KEY (id) );
INSERT INTO cycling.cyclist_stats (id, lastname, basics) VALUES (d4da9bc0-e858-9a45-b664-24561fee4ee1, 'FRAME', { birthday:'1993-06-18',nationality:'New Zealand',weight:null,height:null });
INSERT INTO cycling.cyclist_stats (id, lastname, basics) VALUES (4eea43b1-7155-52ba-13c0-d0f8c98963ef, 'VIGANO', { birthday:'1984-06-12',nationality:'Italy',weight:'67 kg',height:'1.82 m' });
INSERT INTO cycling.cyclist_stats (id, lastname, basics) VALUES (320554de-4142-36b1-9f2a-1b3c3b75cdec, 'TIRALONGO', { birthday:'1977-07-08',nationality:'Italy',weight:'63 kg',height:'1.78 m' });
SELECT * FROM cycling.cyclist_stats;
SELECT * FROM cycling.cyclist_stats WHERE id = 454171da-2622-39b5-1c9a-1f2a4c33fdcc;

-- NEW IN C* 3.6
-- UPDATE AND DELETE single fields in UDTs with only non-collection fields
-- CHANGE "CREATE TABLE IN LAST EXAMPLE TO non-frozen
CREATE TABLE cycling.cyclist_stats ( id UUID, lastname text, basics basic_info, PRIMARY KEY (id) );
-- Now birthday can be updated separate from nationality, weight, and height
UPDATE cycling.cyclist_stats SET basics.birthday = '2000-12-12' WHERE id = 050132da-1960-52c4-8c7a-3d4b9a00eacf;

-- Q12:
-- Find total number of PCS points for a particular cyclist
-- CREATE TABLE WITH PRIMARY KEY: PARTITION KEY + 1 CLUSTERING COLUMN
-- USE STANDARD AGGREGATE IN QUERY
CREATE TABLE cycling.cyclist_points (id UUID, firstname text, lastname text, race_title text, race_points int, PRIMARY KEY (id, race_points ));
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (d8c12df2-367b-2a1c-5a3c-dfbaa2d98aea, 'Giorgia','BRONZINI', 'Tour of Chongming Island World Cup', 120);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (d6e02cd4-161c-1d2f-4f0e-edabc3d17cca, 'Giorgia','BRONZINI', 'Trofeo Alfredo Binda - Comune di Cittiglio', 6);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (c2c70bf4-603f-7f2b-5d0f-daebb4c18add, 'Giorgia','BRONZINI', 'Acht van Westerveld', 75);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (367527fd-0286-74d7-1b5c-5b3c6e33bdef, 'Paolo','TIRALONGO', '98th Giro d''Italia - Stage 15', 2);
SELECT sum(race_points) FROM cycling.cyclist_points WHERE id=c7b72fc2-542f-7e0a-0e4c-fcbea8e10bab;

-- Q13:
-- USES TABLE cycling.cyclist_points
-- Find total number of PCS points for a particular cyclist using a user-defined function (UDF) created using java function log
-- cassandra.yaml must be modified to allow UDFs to work
-- enable_user_defined_functions: true (false by default)
-- CREATE UDF
CREATE TABLE cycling.cyclist_points (id UUID, firstname text, lastname text, race_title text, race_points double, PRIMARY KEY (id, race_points ));
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (f8d29df4-193b-0f5c-7e3b-aeeec7e12dab, 'Giorgia','BRONZINI', 'Tour of Chongming Island World Cup', 120);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (d4c80ed2-060a-8f7d-9a6c-bdcbf7f28ccb, 'Giorgia','BRONZINI', 'Trofeo Alfredo Binda - Comune di Cittiglio', 6);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (c3b29fa1-708c-7d3f-5f8f-eacef3d82eef, 'Giorgia','BRONZINI', 'Acht van Westerveld', 75);
INSERT INTO cycling.cyclist_points (id, firstname, lastname, race_title, race_points) VALUES (365741ed-6400-27b8-5e1f-4e1f4f13fcca, 'Paolo','TIRALONGO', '98th Giro d''Italia - Stage 15', 2);
CREATE OR REPLACE FUNCTION cycling.fLog (input double) CALLED ON NULL INPUT RETURNS double LANGUAGE java AS 'return Double.valueOf(Math.log(input.doubleValue()));';
SELECT id, lastname, fLog(race_points) FROM cycling.cyclist_points;

-- Q14:
--Find the average race_time in seconds for a particular race for a particular team.
-- CREATE UDA that computes the average value
--CREATE TABLE WITH SIMPLE PRIMARY KEY: PARTITION KEY + 2 CLUSTERING COLUMNS
CREATE OR REPLACE FUNCTION cycling.avgState ( state tuple<int,bigint>, val int ) CALLED ON NULL INPUT RETURNS tuple<int,bigint> LANGUAGE java AS 'if (val !=null) { state.setInt(0, state.getInt(0)+1); state.setLong(1, state.getLong(1)+val.intValue()); } return state;';
CREATE OR REPLACE FUNCTION cycling.avgFinal ( state tuple<int,bigint> ) CALLED ON NULL INPUT RETURNS double LANGUAGE java AS 'double r = 0; if (state.getInt(0) == 0) return null; r = state.getLong(1); r/= state.getInt(0); return Double.valueOf(r);';
CREATE AGGREGATE cycling.average ( int ) SFUNC avgState STYPE tuple<int,bigint> FINALFUNC avgFinal INITCOND (0,0);
CREATE TABLE cycling.team_average (team_name text, cyclist_name text, cyclist_time_sec int, race_title text, PRIMARY KEY (team_name, race_title,cyclist_name));
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('UnitedHealthCare Pro Cycling Womens Team','Katie HALL',11449,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('UnitedHealthCare Pro Cycling Womens Team','Linda VILLUMSEN',11485,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('UnitedHealthCare Pro Cycling Womens Team','Hannah BARNES',11490,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('Velocio-SRAM','Alena AMIALIUSIK',11451,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('Velocio-SRAM','Trixi WORRACK',11453,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
INSERT INTO cycling.team_average (team_name, cyclist_name, cyclist_time_sec, race_title) VALUES ('TWENTY16 presented by Sho-Air','Lauren KOMANSKI',11451,'Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe');
SELECT cycling.average(cyclist_time_sec) FROM cycling.team_average WHERE team_name='UnitedHealthCare Pro Cycling Womens Team' AND race_title='Amgen Tour of California Women''s Race presented by SRAM - Stage 1 - Lake Tahoe > Lake Tahoe';

-- Q15:
-- CREATE INDEX - PARTITION KEY
-- Uses cycling.rank_by_year_and_name
-- Find rank for all races for a particular race year
CREATE INDEX ryear ON cycling.rank_by_year_and_name (race_year);
-- This will not work without the index, because the table has a composite partition key
SELECT * FROM cycling.rank_by_year_and_name WHERE race_year=2015;
-- INDEX on clustering column
CREATE INDEX rrank ON cycling.rank_by_year_and_name (rank);
SELECT * FROM cycling.rank_by_year_and_name WHERE rank = 1;

-- Q16:
-- CREATE INDEX - COLLECTION - SET
-- Find all the cyclists that have been on a particular team
CREATE INDEX team ON cycling.cyclist_career_teams (teams);
SELECT * FROM cycling.cyclist_career_teams WHERE teams CONTAINS 'Nederland bloeit';
SELECT * FROM cycling.cyclist_career_teams WHERE teams CONTAINS 'Rabobank-Liv Giant';

-- Q17:
-- CREATE INDEX - COLLECTION ON MAP KEYS
-- Find all cyclist/team combinations for a particular year
-- CREATE TABLE cycling.cyclist_teams ( id UUID PRIMARY KEY, lastname text, firstname text, teams map<int,text> );
CREATE INDEX team_year ON cycling.cyclist_teams (KEYS(teams));
SELECT * FROM cycling.cyclist_teams WHERE teams CONTAINS KEY 2015;

-- Q35:
-- CREATE INDEX - ENTRIES ON MAP KEYS
-- ONLY VALID FOR MAP TYPE
CREATE TABLE cycling.birthday_list (cyclist_name text PRIMARY KEY, blist map<text,text>);
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Allan DAVIS', {'age':'35', 'bday':'27/07/1980', 'nation':'AUSTRALIA'});
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Claudio VANDELLI', {'age':'54', 'bday':'27/07/1961', 'nation':'ITALY'});
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Laurence BOURQUE', {'age':'23', 'bday':'27/07/1992', 'nation':'CANADA'});
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Claudio HEINEN', {'age':'23', 'bday':'27/07/1992', 'nation':'GERMANY'});
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Luc HAGENAARS', {'age':'28', 'bday':'27/07/1987', 'nation':'NETHERLANDS'});
INSERT INTO cycling.birthday_list (cyclist_name, blist) VALUES ('Toine POELS', {'age':'52', 'bday':'27/07/1963', 'nation':'NETHERLANDS'});
CREATE INDEX blist_idx ON cycling.birthday_list (ENTRIES(blist));
SELECT * FROM cycling.birthday_list WHERE blist['age'] = '23';
SELECT * FROM cycling.birthday_list WHERE blist['nation'] = 'GERMANY';
SELECT * FROM cycling.birthday_list WHERE blist['bday'] = '27/07/1992';

-- Q36:
-- CREATE INDEX - FULL ON FROZEN COLLECTION
-- ONLY VALID FOR FROZEN COLLECTIONS (SET, LIST, MAP)
CREATE TABLE cycling.race_starts (cyclist_name text PRIMARY KEY, rnumbers FROZEN<LIST<int>>);
CREATE INDEX rnumbers_idx ON cycling.race_starts (FULL(rnumbers));
INSERT INTO cycling.race_starts (cyclist_name,rnumbers) VALUES ('Alexander KRISTOFF',[40,5,14]);
INSERT INTO cycling.race_starts (cyclist_name,rnumbers) VALUES ('Alejandro VALVERDE',[67,17,20]);
INSERT INTO cycling.race_starts (cyclist_name,rnumbers) VALUES ('Alberto CONTADOR',[61,14,7]);
INSERT INTO cycling.race_starts (cyclist_name,rnumbers) VALUES ('Christopher FROOME',[28,10,6]);
INSERT INTO cycling.race_starts (cyclist_name,rnumbers) VALUES ('John DEGENKOLB',[39,7,14]);
SELECT * FROM cycling.race_starts WHERE rnumbers = [39,7,14];

-- NOT A QUERY, JUST AN EXAMPLE
-- INSERT DATA IN JSON FORMAT
INSERT INTO cycling.cyclist_category JSON '{ "category" : "GC", "points" : 780, "id" : "049cd07b-6cca-688b-d3fd-63500b815fff", "lastname" : "SUTHERLAND" }';
-- null INSERTION EXAMPLE
INSERT INTO cycling.cyclist_category JSON '{ "category" : "Sprint", "points" : 700, "id" : "351df57b-7fdd-376c-f7cb-77172c941fce" }';

-- NOT A QUERY, JUST AN EXAMPLE
-- UPDATE SET
-- Can only be +
-- Add team to a cyclist's list of teams, order doesn't matter; this example adds it to the end
UPDATE cycling.cyclist_career_teams SET teams = teams + {'Team DSB - Ballast Nedam'} WHERE id=5f7690aa-6d68-5d53-8c14-cebfc6a613e2;

-- NOT A QUERY, JUST AN EXAMPLE
-- UPDATE LIST
-- Add events to the events list with either +/- or a specific place in the list like events[2]
UPDATE cycling.upcoming_calendar SET events = ['The Parx Casino Philly Cycling Classic'] + events WHERE year = 2015 AND month = 06;
UPDATE cycling.upcoming_calendar SET events[2] = 'Vuelta Ciclista a Venezuela' WHERE year = 2015 AND month = 06;

-- NOT A QUERY, JUST AN EXAMPLE
-- UPDATE MAP
-- Can only be +
UPDATE cycling.cyclist_teams SET teams = teams + {2009 : 'DSB Bank - Nederland bloeit'} WHERE id = 7a7365df-3b54-2c73-1f52-eefdf1a539f4;
SELECT teams FROM cycling.cyclist_teams WHERE id = 3d0332de-6f25-7c70-2b12-aefdd9b499c2;
UPDATE cycling.cyclist_teams SET teams[2006] = 'Team DSB - Ballast Nedam' WHERE id = 4f8918ca-6f00-7e21-0d30-baaad3f555d7;

-- Q22:
-- UPDATE AND SELECT USING TTL
-- QUERY TO FIND TIME-TO-LIVE
-- Insert is to put in dummy record, UPDATE gives it a TTL
-- Repeated use of the SELECT will show the TTL as it counts down
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (200, 'placeholder', '2015-05-27', '2015-05-27') USING TTL;
UPDATE cycling.calendar USING TTL 300 SET race_name = 'dummy' WHERE race_id = 200 AND race_start_date = '2015-05-27' AND race_end_date = '2015-05-27';
SELECT TTL(race_name) FROM cycling.calendar WHERE race_id=200;

-- Q18:
-- QUERY WITH ORDER BY
-- Find all calendar events for a particular year and order by month
SELECT * FROM cycling.upcoming_calendar WHERE year= 2015 ORDER BY month DESC;

-- Q19:
-- QUERY WITH INEQUALITIES
-- Find all calendar events for a particular year between two set months
SELECT * FROM cycling.upcoming_calendar WHERE year = 2015 AND month <= 06 AND month >= 07;

-- NOT A QUERY, REALLY, JUST AN EXAMPLE
-- SELECT and GET RESULTS in JSON FORMAT
SELECT JSON month, year, events FROM cycling.upcoming_calendar;

-- Q20:
-- QUERY - WHERE ... IN SIMPLE
-- Notice the difference between using 'ORDER BY points DESC' and not using it - changes the order of reporting
-- Find all cyclists for a particular category and order by points
PAGING OFF;
SELECT * FROM cycling.cyclist_category WHERE category IN ('Time-trial', 'Sprint') ORDER BY id DESC;
PAGING OFF;
SELECT * FROM cycling.cyclist_category WHERE category IN ('Time-trial', 'Sprint') ORDER BY id ASC;

-- Q21:
-- QUERY - WHERE ... IN COMPLEX
-- Find particular races in a range of start and end dates
PAGING OFF;
SELECT * FROM cycling.calendar WHERE race_id IN (100, 101, 102) AND (race_start_date, race_end_date) IN (('2015-05-09','2015-05-31'),('2015-05-06', '2015-05-31'));
PAGING OFF;
SELECT * FROM cycling.calendar WHERE race_id IN (100, 101, 102) AND (race_start_date, race_end_date) >= ('2015-05-09','2015-05-24');

-- Q23 and 24:
-- Standard Aggregates
-- Find sum of cyclist points for a particular cyclist
-- Find the number of cyclists from a particular country
SELECT sum(race_points) FROM cycling.cyclist_points WHERE id = e7c85fb8-529d-8a8d-4f6e-afece9a14caa;
SELECT count(cyclist_name) FROM cycling.country_flag WHERE country='Belgium';

-- Q25
-- QUERY - SCAN A PARTITION
-- Find all cyclists that finished a race in a particular window of time
CREATE TABLE cycling.race_times (race_name text, cyclist_name text, race_time text, PRIMARY KEY (race_name, race_time));
INSERT INTO cycling.race_times (race_name, cyclist_name, race_time) VALUES ('17th Santos Tour Down Under', 'Rohan DENNIS', '19:15:18');
INSERT INTO cycling.race_times (race_name, cyclist_name, race_time) VALUES ('17th Santos Tour Down Under', 'Richie PORTE', '19:15:20');
INSERT INTO cycling.race_times (race_name, cyclist_name, race_time) VALUES ('17th Santos Tour Down Under', 'Cadel EVANS', '19:15:38');
INSERT INTO cycling.race_times (race_name, cyclist_name, race_time) VALUES ('17th Santos Tour Down Under', 'Tom DUMOULIN', '19:15:40');
SELECT * FROM cycling.race_times WHERE race_name = '17th Santos Tour Down Under' AND race_time >= '19:15:19' AND race_time <= '19:15:39';

-- NOT A QUERY, JUST AN EXAMPLE:
-- BATCH statement
-- Insert data into multiple tables using a BATCH statement
-- Note that what is inserted is data for the SAME cyclist, to two tables
BEGIN BATCH
  INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (b5badfb1-d216-9712-4220-e50b5043ab5a, 'PIETERS', 'Amy');
  INSERT INTO cycling.cyclist_id (lastname, firstname, age, id) VALUES ('PIETERS', 'Amy', 23, c3fdfea6-e030-6957-9867-a30d0635ad2d);
APPLY BATCH;
SELECT * FROM cycling.cyclist_name;
SELECT * FROM cycling.cyclist_id;

-- NOT A QUERY, JUST AN EXAMPLE:
-- BATCH statement MISUSE
-- Insert data into same table, but involves multiple nodes due to partition key = id
BEGIN BATCH
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES  (5a8f4826-98a1-45bb-6bbe-92b575d32098,'HOSKINS', 'Melissa');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES  (84bd66f0-60de-7ee9-cd76-a702be967830,'FERNANDES', 'Marcia');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES  (6295e3de-a53b-1d4b-72c0-f3a995139ae2,'NIEWIADOMA', 'Katarzyna');
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES  (84ceca2c-039a-4ed4-b5f0-107b05b89179,'ADRIAN', 'Vera');
APPLY BATCH;

-- NOT A QUERY, JUST AN EXAMPLE:
-- BATCH statement WITH CONDITIONAL "IF NOT EXISTS"
-- EXAMPLE USES CYCLIST'S EXPENSES
CREATE TABLE cycling.cyclist_expenses ( cyclist_name text, balance float STATIC, expense_id int, amount float, description text, paid boolean, PRIMARY KEY (cyclist_name, expense_id) );
BEGIN BATCH
INSERT INTO cycling.cyclist_expenses (cyclist_name, balance) VALUES ('Vera ADRIAN', 0) IF NOT EXISTS;
INSERT INTO cycling.cyclist_expenses (cyclist_name, expense_id, amount, description, paid) VALUES ('Vera ADRIAN', 1, 7.95, 'Breakfast', false);
APPLY BATCH;

UPDATE cycling.cyclist_expenses SET balance = -7.95 WHERE cyclist_name = 'Vera ADRIAN' IF balance = 0;

-- NOT A QUERY, JUST AN EXAMPLE:
-- BATCH statement WITH CONDITIONAL "IF"
BEGIN BATCH
INSERT INTO cycling.cyclist_expenses (cyclist_name, expense_id, amount, description, paid) VALUES ('Vera ADRIAN', 2, 13.44, 'Lunch', true);
INSERT INTO cycling.cyclist_expenses (cyclist_name, expense_id, amount, description, paid) VALUES ('Vera ADRIAN', 3, 25.00, 'Dinner', false);
UPDATE cycling.cyclist_expenses SET balance = -32.95 WHERE cyclist_name = 'Vera ADRIAN' IF balance = -7.95;
APPLY BATCH;

-- NOT A QUERY, JUST AN EXAMPLE:
-- BATCH statement WITH CONDITIONAL "IF"
BEGIN BATCH
UPDATE cycling.cyclist_expenses SET balance = 0 WHERE cyclist_name = 'Vera ADRIAN' IF balance = -32.95;
UPDATE cycling.cyclist_expenses SET paid = true WHERE cyclist_name = 'Vera ADRIAN' AND expense_id = 1 IF paid = false;
UPDATE cycling.cyclist_expenses SET paid = true WHERE cyclist_name = 'Vera ADRIAN' AND expense_id = 3 IF paid = false;
APPLY BATCH;

-- NOT A QUERY, JUST AN EXAMPLE
-- LIGHTWEIGHT TRANSACTION
-- Insert or update information using a conditional statement
INSERT INTO cycling.cyclist_name (id, lastname, firstname) VALUES (c9b58060-bb37-6044-66b2-b6c3f15f794c, 'RATTO', 'Rissella') IF NOT EXISTS;

-- UPDATE USING LIGHTWEIGHT TRANSACTION
UPDATE cycling.cyclist_name SET firstname = 'Rossella' WHERE id=e7c99655-eb00-7326-71f4-e9f9c38a574e IF lastname = 'RATTO';

-- Q26
-- QUERY USING MULTIPLE INDEXES
-- DISCUSSION OF THE NEED FOR ALLOW FILTERING
-- IS THIS BETTER THAN cyclist_stats??
CREATE TABLE cycling.cyclist_alt_stats ( id UUID PRIMARY KEY, lastname text, birthday timestamp, nationality text, weight text, height text );
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (cd521c27-89f2-9c99-7a55-4df1886f7003,'TSATEVICH', '1989-07-05', 'Russia', '64 kg', '1.69 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (d8e17094-1ff1-45a1-2bb4-523364f74ffb,'ISAYCHEV', '1986-04-21', 'Russia', '80 kg', '1.88 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (744bd191-8e58-44c6-ab83-2868f7a4323d,'BELKOV', '1985-01-09', 'Russia', '71 kg', '1.84 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (a3888271-79cb-5c17-2d33-5c6781570565,'BRUTT', '1982-01-29', 'Russia', '68 kg', '1.78 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (923950d0-30aa-0038-cd09-becfb54da208,'LAGUTIN', '1981-01-14', 'Russia', '63 kg', '1.82 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (a55a1c83-4037-8af2-6974-a3548b51c196,'GRMAY', '1991-08-25', 'Ethiopia', '63 kg', '1.75 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (d75c8858-22ff-229a-8938-c9aad5a503e4,'VEIKKANEN', '1981-03-29', 'Finland', '66 kg', '1.78 m');
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (c5ebbd66-7b10-3767-b03d-b81ce0da99a5,'TLEUBAYEV', '1987-03-07', 'Kazakhstan', null, null);
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (8bb2736a-31ba-2848-b070-cb8ef766ec5a,'PAULINHO', '1990-05-27', 'Portugal', null, null);
INSERT INTO cycling.cyclist_alt_stats (id, lastname, birthday, nationality, weight, height) VALUES (0ffe790b-45cc-8d50-25b3-17951147dd93,'DUVAL', '1990-05-27','France', null, null);
CREATE INDEX birthday_idx ON cycling.cyclist_alt_stats (birthday);
CREATE INDEX nationality_idx ON cycling.cyclist_alt_stats (nationality);
SELECT * FROM cycling.cyclist_alt_stats WHERE birthday = '1982-01-29' AND nationality = 'Russia' ALLOW FILTERING;
SELECT * FROM cycling.cyclist_alt_stats WHERE birthday = '1990-05-27' AND nationality = 'Portugal' ALLOW FILTERING;

-- Q27
-- USING EXPIRING DATA AND TTL TO DISPLAY THE LAST 3 DAYS race data
-- 3 days in seconds is 259,200
-- 2 days in seconds is 172800
-- Data will vanish when its TTL runs out
CREATE TABLE cycling.last_3_days (race_name text, year timestamp, rank int, cyclist_name text, PRIMARY KEY (year, rank, cyclist_name));
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('Giro d''Italia Stage 16','2015-05-26',1,'Mikel Landa') USING TTL 259200;
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('Giro d''Italia Stage 16','2015-05-26',2,'Steven Kruijswijk') USING TTL 259200;
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('Giro d''Italia Stage 16','2015-05-26',3,'Alberto Contador') USING TTL 259200;
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('National Championships United States - Road Race (NC)','2015-05-25',1,'Matthew Busche') USING TTL 172800;
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('National Championships United States - Road Race (NC)','2015-05-25',2,'Joe Dombrowski') USING TTL 172800;
INSERT INTO cycling.last_3_days (race_name, year, rank, cyclist_name) VALUES ('National Championships United States - Road Race (NC)','2015-05-25',3,'Kiel Reijnen') USING TTL 172800;
SELECT TTL(race_name) FROM cycling.last_3_days;
SELECT TTL(race_name) FROM cycling.last_3_days;
SELECT * FROM cycling.last_3_days; // WILL ONLY SHOW NON-EXPIRED ROWS

-- Q28:
-- QUERY USING FUNCTION TOKEN()
-- Note how results are not consistent with dates alone; partitioner order is how they are returned
-- All 6 entries show
SELECT * FROM cycling.last_3_days WHERE token(year) > token ('2015-05-24');
-- No entries show
SELECT * FROM cycling.last_3_days WHERE token(year) > token ('2015-05-25');
-- 3 entries for 2015-05-25 show
SELECT * FROM cycling.last_3_days WHERE token(year) > token ('2015-05-26');
-- No entries show
SELECT * FROM cycling.last_3_days WHERE token(year) > token ('2015-05-27');
SELECT token(year) FROM cycling.last_3_days; //PRINTS partition hash
-- MIXED TOKEN AND PARTITION KEY
SELECT * FROM cycling.last_3_days WHERE token(year) < token ('2015-05-26') AND year IN ('2015-05-24','2015-05-25');


-- DELETE WHOLE ROW
-- Leave column(s) blank
DELETE FROM cycling.calendar WHERE race_id = 200;

-- DELETE COLUMN VALUE
DELETE lastname FROM cycling.cyclist_name WHERE id = b7cabbf1-c543-8877-4718-f45b6232fb0f;
UPDATE cycling.cyclist_name SET lastname = 'PIETERS' WHERE id = c9debda5-d908-9537-0950-b61b1253eb9c; // TO RESTORE THE COLUMN VALUE

-- DELETE ITEM FROM LIST
DELETE events[2] FROM cycling.upcoming_calendar WHERE year = 2015 AND month = 06;

-- DELETE ITEM FROM MAP
DELETE teams[2009] FROM cycling.cyclist_teams WHERE id=b1fd0153-fe7b-1284-e71d-1084fdc0dcff;
UPDATE cycling.cyclist_teams SET teams = teams + {2009 : 'Team Flexpoint' } WHERE id = e0ee1128-ca3d-5691-a45f-7282fbe5ddba; // TO RESTORE THE MAP VALUE

-- ALTER TABLE
-- ADD COLUMN
ALTER TABLE cycling.cyclist_alt_stats ADD age int;

-- ALTER TABLE WITH COLLECTION
ALTER TABLE cycling.upcoming_calendar ADD description map<text,text>;
UPDATE cycling.upcoming_calendar SET description = description + {'Criterium du Dauphine' : 'Easy race', 'Tour du Suisse' : 'Hard uphill race'} WHERE year = 2015 AND month = 6;

-- ALTER TABLE AND ALTER COLUMN TYPE
-- ADDS COLUMN as varchar and then changes it to text
ALTER TABLE cycling.cyclist_alt_stats ADD favorite_color varchar;
ALTER TABLE cycling.cyclist_alt_stats ALTER favorite_color TYPE text;


-- ALTER TYPE
ALTER TYPE cycling.fullname ADD middlename text;
ALTER TYPE cycling.fullname RENAME middlename TO middleinitial;

-- TUPLE WAS USED IN THE UDA TO HOLD 2 values - see example in UDA section

-- Q29:
-- TUPLE
-- Store the latitude/longitude waypoints for the route of a race
CREATE TABLE cycling.route (race_id int, race_name text, point_id int, lat_long tuple<text, tuple<float,float>>, PRIMARY KEY (race_id, point_id));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 1, ('Onnens', (46.8444,6.6667)));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 2, ('Champagne', (46.833, 6.65)));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 3, ('Novalle', (46.833, 6.6)));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 4, ('Vuiteboeuf', (46.8, 6.55)));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 5, ('Baulmes', (46.7833, 6.5333)));
INSERT INTO cycling.route (race_id, race_name, point_id, lat_long) VALUES (500, '47th Tour du Pays de Vaud', 6, ('Les Clées', (46.7222, 6.5222)));
SELECT race_name, point_id, lat_long AS City_Latitude_Longitude FROM cycling.route; // Showcases 'AS' to rename column header

-- Q30:
-- QUERY USING DISTINCT
-- Find all the distinct race_id values from cycling.route
SELECT DISTINCT race_id from cycling.route;

-- Q31:
-- TUPLE
-- Rank nations by points, including top cyclist
-- tuple is rank, name, points
CREATE TABLE cycling.nation_rank ( nation text PRIMARY KEY, info tuple<int,text,int> );
INSERT INTO cycling.nation_rank (nation, info) VALUES ('Spain', (1,'Alejandro VALVERDE' , 9054));
INSERT INTO cycling.nation_rank (nation, info) VALUES ('France', (2,'Sylvain CHAVANEL' , 6339));
INSERT INTO cycling.nation_rank (nation, info) VALUES ('Belgium', (3,'Phillippe GILBERT' , 6222));
INSERT INTO cycling.nation_rank (nation, info) VALUES ('Italy', (4,'Davide REBELLINI' , 6090));
SELECT * FROM cycling.nation_rank;

-- Q32:
-- TUPLE
-- Popular Riders
CREATE TABLE cycling.popular (rank int PRIMARY KEY, cinfo tuple<text,text,int> );
INSERT INTO cycling.popular (rank, cinfo) VALUES (1, ('Spain', 'Mikel LANDA', 1137));
INSERT INTO cycling.popular (rank, cinfo) VALUES (2, ('Netherlands', 'Steven KRUIJSWIJK', 621));
INSERT INTO cycling.popular (rank, cinfo) VALUES (3, ('USA', 'Matthew BUSCHE', 230));
INSERT INTO cycling.popular (rank, cinfo) VALUES (4, ('Italy', 'Fabio ARU', 163));
INSERT INTO cycling.popular (rank, cinfo) VALUES (5, ('Canada', 'Ryder HESJEDAL', 148));
SELECT * FROM cycling.popular;

-- Q33:
-- COUNTER TABLE
-- Keep the count for popularity, incrementing or decrementing
CREATE TABLE cycling.popular_count ( id UUID PRIMARY KEY, popularity counter );
UPDATE cycling.popular_count SET popularity = popularity + 1 WHERE id = 0fc63eaa-f48a-47e2-f3e3-41e8bd3c2e35;
SELECT * FROM cycling.popular_count;
UPDATE cycling.popular_count SET popularity = popularity + 125 WHERE id = 7ae25cbc-d35e-08c1-b8a7-03b7fd4b5d31;
SELECT * FROM cycling.popular_count;
UPDATE cycling.popular_count SET popularity = popularity - 64 WHERE id = 1cd05daf-e54b-33f5-c3f9-04d3fd1a4f82;
SELECT * FROM cycling.popular_count;

-- Q34:
-- Find the writetime for a column in a table
SELECT WRITETIME (firstname) FROM cycling.cyclist_points WHERE id=183091ec-2274-74f4-7f8e-6b1e4f26cafd;

-- NOT A QUERY
-- INSERTING STRING CONSTANT USING DOUBLE DOLLAR SIGNS
INSERT INTO cycling.calendar (race_id, race_start_date, race_end_date, race_name) VALUES (201, '2015-02-18', '2015-02-22', $$Women's Tour of New Zealand$$);

-- ROLES, USERS, PERMISSIONS
-- cassandra.yaml must be changed to allow login with username and password
-- authenticator: PasswordAuthenticator (AllowAllAuthenticator by default)
-- authorizer: CassandraAuthorizer (AllowAllAuthorizer by default)
CREATE USER IF NOT EXISTS sandy WITH PASSWORD 'Wwhw4Jjm@' NOSUPERUSER;
CREATE USER chuck WITH PASSWORD 'Xlqehk5ur$' SUPERUSER;
ALTER USER sandy SUPERUSER;
LIST USERS;
-- DROP USER IF EXISTS chuck;
CREATE ROLE IF NOT EXISTS team_manager WITH PASSWORD = 'QqcwVw1Ml!';
CREATE ROLE sys_admin WITH PASSWORD = 'BtrwCaHx3ai' AND LOGIN = true AND SUPERUSER = true;
ALTER ROLE sys_admin WITH PASSWORD = 'Rfb0cmu5quaQzn' AND SUPERUSER = false;
GRANT sys_admin TO team_manager;
GRANT team_manager TO sandy;
LIST ROLES;
LIST ROLES OF sandy;
REVOKE sys_admin FROM team_manager;
REVOKE team_manager FROM sandy;
DROP ROLE IF EXISTS sys_admin;
GRANT MODIFY ON KEYSPACE cycling TO team_manager;
GRANT DESCRIBE ON ALL ROLES TO sys_admin;
GRANT AUTHORIZE ALL KEYSPACES TO sys_admin;
REVOKE SELECT ON ALL KEYSPACES FROM team_manager;
REVOKE EXECUTE ON FUNCTION cycling.fLog(double) FROM team_manager;
LIST ALL PERMISSIONS OF sandy;
LIST ALL PERMISSIONS ON cycling.cyclist_name OF chuck;

-- Q35:
-- MATERIALIZED VIEW
CREATE TABLE cycling.cyclist_mv (cid UUID PRIMARY KEY, name text, age int, birthday date, country text);
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (f3fc1bd4-c489-9a36-e664-08592cef4dd4,'Alex FRAME', 22, 1993-06-18, 'New Zealand');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (813075ae-5077-16f3-5b2c-2a7b5f16cceb,'Paolo TIRALONGO', 38, '1977-07-08', 'Italy');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (2ff63cba-d25a-39d5-d1b9-22b1dd5d2c74,'Steven KRUIKSWIJK', 28, '1987-06-07', 'Netherlands');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (effeb5f4-9be9-53e6-eccc-0fdaaaed2715,'Pascal EENKHOORN', 18, '1997-02-08', 'Netherlands');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (87d105ef-c690-6be5-e5e2-c4b5ce4cc106,'Bram WELTEN', 18, '1997-03-29', 'Netherlands');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (91f665da-f692-2ba7-aa0b-2d9522396411,'Adrien COSTA', 18, '1997-08-19', 'United States');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (858bf43c-90c6-5f3f-606d-a482dfe3848c,'Joakim BUKDAL', 20, '1994-09-04', 'Denmark');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (d6a5d301-5c6c-0386-5403-2e170c58f5e4,'Cristian EGIDIO', 27, '1987-09-04', 'Brazil');
INSERT INTO cycling.cyclist_mv (cid,name,age,birthday,country) VALUES (d2fba71d-cb66-18e0-dd0b-616c4ee2a44e,'Johannes HEIDER', 27, '1987-09-04','Germany');
CREATE MATERIALIZED VIEW cycling.cyclist_by_age AS SELECT age, birthday, name, country FROM cyclist_mv WHERE age is NOT NULL AND cid IS NOT NULL PRIMARY KEY (age, cid);
CREATE MATERIALIZED VIEW cycling.cyclist_by_country AS SELECT age, birthday, name, country FROM cyclist_mv WHERE country is NOT NULL AND cid IS NOT NULL PRIMARY KEY (country, cid);
CREATE MATERIALIZED VIEW cycling.cyclist_by_birthday AS SELECT age, birthday, name, country FROM cyclist_mv WHERE birthday is NOT NULL AND cid IS NOT NULL PRIMARY KEY (birthday, cid);
--DROP MATERIALIZED VIEW cycling.cyclist_by_age;

-- Q36:
-- USING TIMESTAMP
INSERT INTO cycling.calendar (race_id, race_name, race_start_date, race_end_date) VALUES (200, 'placeholder', '2015-05-27', '2015-05-27') USING TIMESTAMP 123456789;

-- exit
\q
