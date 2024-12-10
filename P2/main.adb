with Ada.Text_IO;
use Ada.Text_IO;
with Ada.Strings.Unbounded;
use Ada.Strings.Unbounded;
with def_maquina;
use def_maquina;
with Gen_Random;

procedure Main is
    -- Inicialitza els noms dels clients
    type Noms is array(1..10) of Unbounded_String;  
    noms_clients : Noms := ( 
        To_Unbounded_String("Aina"),
        To_Unbounded_String("Bernat"),
        To_Unbounded_String("Bel"),
        To_Unbounded_String("Albert"),
        To_Unbounded_String("Laura"),
        To_Unbounded_String("Dídac"),
        To_Unbounded_String("Pep"),
        To_Unbounded_String("Miquel"),
        To_Unbounded_String("Caterina"),
        To_Unbounded_String("Jaume"));
    -- Inicialitza els valors aleatoris de clients i reposadors
    num_clients : Integer := Gen_Random.RandomNumber;
    num_reposadors : Integer := Gen_Random.RandomNumber;
    -- Variables de temps per a la simulació
    temps_client : Duration := 0.1;
    temps_reposador : Duration := 0.3;
    -- Definició de la màquina de refrescs que serà un monitor
    maquina : def_maquina.Maquina;
    -- Definició dels clients
    task type Client is
        entry Entrar(nom_client : Unbounded_String; temps_client : Duration);
    end Client;
    -- Implementació dels clients
    task body Client is
        nom : Unbounded_String;
        temps : Duration;
        num_consumicions : Integer;
        consumicio : Integer := 0;
    begin
        -- El client entra a la màquina i se li assignen els valors corresponents
        accept Entrar(nom_client : Unbounded_String; temps_client : Duration) do
            nom := nom_client;
            temps := temps_client;
        end Entrar;
        -- El client decideix aleatòriament quants refrescs vol
        num_consumicions := Gen_Random.RandomNumber;
        delay temps;
        -- Si no hi ha reposadors a la màquina, el client se'n va
        if (num_reposadors = 0) then
            Put_Line(To_String(nom) & " diu: No hi ha reposadors a la màquina, me'n vaig");
        -- Si hi ha reposadors a la màquina, el client es queda
        else
            Put_Line(To_String(nom) & " diu: Hola, avui faré" & Integer'Image(num_consumicions) & " consumicions");
            -- El client consumeix els refrescs que ha decidit que anava a consumir
            for i in 1 .. num_consumicions loop
                consumicio := consumicio + 1;
                maquina.consumint(To_String(nom), consumicio, num_consumicions);
                delay temps;
            end loop;
            -- El client informa de que ha acabat les seves consumicions i se'n va
            num_clients := num_clients - 1;
            Put_Line(To_String(nom) & " acaba i se'n va, queden" & num_clients'Img & " clients >>>>>>>>>>");
        end if;
    end Client;
    -- Definició dels reposadors
    task type Reposador is
        entry Entrar(id_reposador : Integer; temps_reposador : Duration);
    end Reposador;
    -- Implementació dels reposadors
    task body Reposador is
        id : Integer;
        temps : Duration;
    begin
        -- El reposador entra a la màquina i se li assignen els valors corresponents
        accept Entrar(id_reposador : Integer; temps_reposador : Duration) do
            id := id_reposador;
            temps := temps_reposador;
        end Entrar;
        Put_Line("     El reposador" & id'Img & " comença a treballar");
        -- El reposador reposa els refrescs de la màquina fins que no quedin clients
        while not (num_clients = 0) loop
            delay temps;
            maquina.reposant(id);
        end loop;
        -- El reposador informa de que acaba i se'n va
        Put_Line("++++++++++ El reposador" & id'Img & " diu: No hi ha clients me'n vaig");
        delay temps;
        Put_Line("     El reposador" & id'Img & " acaba i se'n va >>>>>>>>>>");
    end Reposador;
    -- Creació dels arrays de clients i reposadors
    type Array_Clients is array (1..num_clients) of Client;
    type Array_Reposadors is array (1..num_reposadors) of Reposador;
    clients : Array_Clients;
    reposadors : Array_Reposadors;
    -- Codi principal on realitzem els bucles per a la creació de clients i reposadors
    begin
        maquina.preparar;
        for num in 1..num_clients loop
            clients(num).Entrar(noms_clients(num), temps_client);
        end loop;
        for num in 1..num_reposadors loop
            reposadors(num).Entrar(num, temps_reposador);
        end loop;
    end Main;