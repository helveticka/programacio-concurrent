with Text_IO;
use Text_IO;

package body Def_Maquina is
    protected body Maquina is
        -- Inicialització de la màquina
        procedure preparar is
        begin
            Put_Line("********** La màquina està preparada");
        end preparar;
        -- Procés de consumició de refrescs (només un client pot consumir a la vegada)
        entry consumint(Nom : String; Consumicio : Integer; Num_consumicions : Integer) when not maquina_ocupada and (Num_refrescs > 0) is
        begin
            maquina_ocupada := True;
            num_refrescs := num_refrescs - 1;
            Put_Line("---------- " & Nom & " agafa el refresc" & Integer'Image(Consumicio) & " /" & Integer'Image(num_consumicions) & " a la màquina en queden" & Integer'Image(num_refrescs));
            maquina_ocupada := False;
        end consumint;
        -- Procés de reposició de refrescs (només un reposador pot reposar a la vegada)
        entry reposant(Identif : Integer) when not maquina_ocupada is
        begin
            maquina_ocupada := True;
            if not (num_refrescs = 10) then
                Put_Line("++++++++++ El reposador" & Integer'Image(Identif) & " reposa" & Integer'Image(maxim_refrescs - num_refrescs) & " refrescs, ara n'hi ha" & Integer'Image(maxim_refrescs));
            end if;
            num_refrescs := maxim_refrescs;
            maquina_ocupada := False;
        end reposant;
    end Maquina;
end Def_Maquina;