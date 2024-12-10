import threading
import time
import random
# Declaració de semàfors i mutex
SemaforNoel = threading.Semaphore(0)
SemaforRen = threading.Semaphore(0)
SemaforElfH = threading.Semaphore(3)
SemaforElf = threading.Semaphore(0)
Mutex = threading.Semaphore(0)
# Declaració de variables globals
Ren = 0
IndexRen = 0
Elf = 0
IndexElf = 0
# Declaració de constants
TOTAL_RENS = 9
TOTAL_ELFS = 6
AJUDA_ELFS = 3
NUMJOGUINES = 3
AIXECAR = 18
AJUDAR = 2
# Noms dels rens i dels elfs
NomsRens = ["RUDOLPH", "BLITZER", "DONDER", "CUPID", "COMET", "VIXEN", "PRANCER", "DANCER", "DASHER"]
NomsElfs = ["Taleasin", "Halafarin", "Ailduin", "Adamar", "Galather", "Estelar"]
# Funció de Santa Claus
def santa():
    global Ren
    global Elf
    # El pare Noel comença dormint
    print("-------> El Pare Noel diu: Estic despert però me'n torn a jeure")
    # Bucle
    while True:
        # El pare Noel espera a ser cridat
        SemaforNoel.acquire()
        # Si hi ha 3 elfs amb dubtes
        if Elf == AJUDA_ELFS:
            # Es desperta i atén els elfs
            Elf = 0
            print("-------> El Pare Noel diu: Atendré els dubtes d'aquests 3")
            # Allibera els semàfors dels elfs
            for i in range(AJUDA_ELFS):
                SemaforElfH.release()
            print("-------> El Pare Noel diu: Estic cansat me'n torn a jeure")
            # Allibera el semàfor dels elfs
            for i in range(AJUDA_ELFS):
                SemaforElf.release()
        # Si hi ha 9 rens
        elif Ren == TOTAL_RENS:
            # Es desperta i enganxa els rens
            Ren = 0
            print("-------> Pare Noel diu: Enganxaré els rens i partiré")
            # Allibera els semàfors dels rens
            for i in range(TOTAL_RENS):
                SemaforRen.release()
            Mutex.acquire()
            print("-------> El Pare Noel ha enganxat els rens, ha carregat les joguines i se'n va")
            Mutex.release()
            break
# Funció dels rens
def reindeer():
    global Ren
    global IndexRen
    num = Ren
    Ren += 1
    # El ren se'n va a pasturar
    print("         {} se'n va a pasturar".format(NomsRens[num]))
    time.sleep(random.randint(9, 11))
    IndexRen += 1
    # Si és l'últim ren
    if IndexRen == 9:
        print("         {} diu: Som el darrer en voler podem partir".format(NomsRens[num]))
        # Allibera el semàfor del pare Noel
        SemaforNoel.release()
    # Si no és l'últim ren
    else:
        print("         El ren {} arriba, {}".format(NomsRens[num], IndexRen))
    # El ren espera a ser enganxat al trineu
    SemaforRen.acquire()
    print("         El ren {} està enganxat al trineu".format(NomsRens[num]))
    time.sleep(5)
    Mutex.release()
# Funció dels elfs
def elf():
    global IndexElf
    global Elf
    num = IndexElf
    IndexElf += 1
    # L'elf decideix les joguines que construirà
    print("Hola som l'elf {} construiré {} joguines".format(NomsElfs[num], NUMJOGUINES))
    # Bucles dels elfs
    for i in range(AJUDAR):
        time.sleep(random.randint(1, 5))
        SemaforElfH.acquire()
        elf = Elf + 1
        Elf += 1
        # L'elf demana ajuda al pare Noel
        if elf < 3:
            print("{} diu: tinc dubtes amb la joguina {}".format(NomsElfs[num], i + 1))
        # L'elf desperta al pare Noel
        elif elf == AJUDA_ELFS:
            print("{} diu: Som 3 que tenim dubtes, PARE NOEEEEEL!".format(NomsElfs[num]))
            SemaforNoel.release()
        # L'elf construeix la joguina
        SemaforElf.acquire()
        print("{} diu: Construeixo la joguina amb ajuda".format(NomsElfs[num]))
    print("L'elf {} ha fet les seves joguines i acaba <---------".format(NomsElfs[num]))
# Funció principal
def main():
    threads = []
    s = threading.Thread(target=santa)
    threads.append(s)
    # Creació dels threads dels elfs
    for i in range(TOTAL_ELFS):
        e = threading.Thread(target=elf)
        threads.append(e)
    # Creació dels threads dels rens
    for i in range(TOTAL_RENS):
        r = threading.Thread(target=reindeer)
        threads.append(r)
    # Inicialització dels threads
    for t in threads:
        t.start()
    # Join dels threads
    for t in threads:
        t.join()    
    # Finalització de la simulació
    print("SIMULACIÓ ACABADA")
# Inicialització del programa
if __name__ == "__main__":
    main()