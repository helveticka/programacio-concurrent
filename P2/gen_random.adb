-- Cos del paquet que gestiona la generació de nombres aleatoris
with Ada.Numerics.Discrete_Random;

package body Gen_Random is
   function RandomNumber return Integer is
      subtype Range_Type is Integer range 0 .. 10;
      package Random_Integer is new Ada.Numerics.Discrete_Random(Range_Type);
      use Random_Integer;
      Generator : Random_Integer.Generator;
      Random_Value : Range_Type;
   begin
      Reset(Generator);
      Random_Value := Random(Generator);
      return Random_Value;
   end RandomNumber;
end Gen_Random;