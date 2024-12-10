package Def_Maquina is
    protected type Maquina is
        -- Processos de la màquina de refrescs
        procedure preparar;
        entry consumint(Nom : String; Consumicio : Integer; Num_consumicions : Integer);
        entry reposant(Identif : Integer);
    private
        -- Variables de la màquina de refrescs
        maquina_ocupada : Boolean := False;
        num_refrescs : Integer := 0;
        maxim_refrescs : Integer := 10;
    end Maquina;
end def_maquina;