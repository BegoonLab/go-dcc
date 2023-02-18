EESchema Schematic File Version 4
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title ""
Date ""
Rev ""
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Device:R R1
U 1 1 617C5062
P 5225 3075
F 0 "R1" H 5295 3121 50  0000 L CNN
F 1 "3.3k" H 5295 3030 50  0000 L CNN
F 2 "Resistor_THT:R_Axial_DIN0204_L3.6mm_D1.6mm_P1.90mm_Vertical" V 5155 3075 50  0001 C CNN
F 3 "~" H 5225 3075 50  0001 C CNN
F 4 "R" H 5225 3075 50  0001 C CNN "Spice_Primitive"
F 5 "3.3k" H 5225 3075 50  0001 C CNN "Spice_Model"
F 6 "Y" H 5225 3075 50  0001 C CNN "Spice_Netlist_Enabled"
	1    5225 3075
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 617C6026
P 2575 3075
F 0 "C1" H 2690 3121 50  0000 L CNN
F 1 "10n" H 2690 3030 50  0000 L CNN
F 2 "Capacitor_THT:C_Disc_D3.0mm_W1.6mm_P2.50mm" H 2613 2925 50  0001 C CNN
F 3 "~" H 2575 3075 50  0001 C CNN
F 4 "C" H 2575 3075 50  0001 C CNN "Spice_Primitive"
F 5 "10n" H 2575 3075 50  0001 C CNN "Spice_Model"
F 6 "Y" H 2575 3075 50  0001 C CNN "Spice_Netlist_Enabled"
	1    2575 3075
	1    0    0    -1  
$EndComp
$Comp
L Device:R R2
U 1 1 617C6B7D
P 2125 2625
F 0 "R2" V 2332 2625 50  0000 C CNN
F 1 "10k" V 2241 2625 50  0000 C CNN
F 2 "Resistor_THT:R_Axial_DIN0204_L3.6mm_D1.6mm_P1.90mm_Vertical" V 2055 2625 50  0001 C CNN
F 3 "~" H 2125 2625 50  0001 C CNN
F 4 "R" H 2125 2625 50  0001 C CNN "Spice_Primitive"
F 5 "10k" H 2125 2625 50  0001 C CNN "Spice_Model"
F 6 "Y" H 2125 2625 50  0001 C CNN "Spice_Netlist_Enabled"
	1    2125 2625
	0    -1   -1   0   
$EndComp
Wire Wire Line
	1975 2625 1725 2625
$Comp
L power:GND #PWR0101
U 1 1 617CF5AA
P 3875 3425
F 0 "#PWR0101" H 3875 3175 50  0001 C CNN
F 1 "GND" H 3880 3252 50  0000 C CNN
F 2 "" H 3875 3425 50  0001 C CNN
F 3 "" H 3875 3425 50  0001 C CNN
	1    3875 3425
	1    0    0    -1  
$EndComp
Wire Wire Line
	3875 3425 3875 3225
$Comp
L Driver_Motor:LMD18200 U1
U 1 1 617C7926
P 3875 2325
F 0 "U1" H 3875 3306 50  0000 C CNN
F 1 "LMD18200" H 3875 3215 50  0000 C CNN
F 2 "Package_TO_SOT_THT:TO-220-11_P3.4x5.08mm_StaggerOdd_Lead4.85mm_Vertical" H 2425 975 50  0001 L CNN
F 3 "http://www.ti.com/lit/ds/symlink/lmd18200.pdf" H 3775 2325 50  0001 C CNN
	1    3875 2325
	1    0    0    -1  
$EndComp
Wire Wire Line
	4575 2025 5225 2025
Connection ~ 3875 3225
Wire Wire Line
	5225 3225 3875 3225
Wire Wire Line
	2275 2625 2325 2625
$Comp
L Device:CP C3
U 1 1 617FDA35
P 2075 3075
F 0 "C3" H 2193 3121 50  0000 L CNN
F 1 "4.7m" H 2193 3030 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D22.0mm_P10.00mm_SnapIn" H 2113 2925 50  0001 C CNN
F 3 "~" H 2075 3075 50  0001 C CNN
F 4 "C" H 2075 3075 50  0001 C CNN "Spice_Primitive"
F 5 "4.7m" H 2075 3075 50  0001 C CNN "Spice_Model"
F 6 "Y" H 2075 3075 50  0001 C CNN "Spice_Netlist_Enabled"
	1    2075 3075
	1    0    0    -1  
$EndComp
$Comp
L Device:CP C2
U 1 1 617FE4F3
P 1675 3075
F 0 "C2" H 1793 3121 50  0000 L CNN
F 1 "4.7m" H 1793 3030 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D22.0mm_P10.00mm_SnapIn" H 1713 2925 50  0001 C CNN
F 3 "~" H 1675 3075 50  0001 C CNN
F 4 "C" H 1675 3075 50  0001 C CNN "Spice_Primitive"
F 5 "4.7m" H 1675 3075 50  0001 C CNN "Spice_Model"
F 6 "Y" H 1675 3075 50  0001 C CNN "Spice_Netlist_Enabled"
	1    1675 3075
	1    0    0    -1  
$EndComp
Wire Wire Line
	1675 3225 2075 3225
Connection ~ 2075 3225
Wire Wire Line
	2075 3225 2575 3225
Connection ~ 2575 3225
Wire Wire Line
	2575 3225 3875 3225
Wire Wire Line
	1675 2925 2075 2925
Wire Wire Line
	2075 2925 2575 2925
Connection ~ 2075 2925
Connection ~ 2325 2625
Wire Wire Line
	2325 2625 3175 2625
Wire Wire Line
	2325 2625 2325 2175
Wire Wire Line
	1775 2175 2325 2175
Wire Wire Line
	3125 2025 3175 2025
Text Notes 5675 2175 0    50   ~ 0
К рельсам
Text Notes 1025 2575 0    50   ~ 0
3V3 на Малину
Text Notes 1050 2100 0    50   ~ 0
GPIO27 на малину
Text Notes 2650 1950 0    50   ~ 0
5V на малину
Text Notes 2650 2250 0    50   ~ 0
GPIO17\nна малину
Connection ~ 2575 2925
Text Notes 2275 1475 0    50   ~ 0
+18V
$Comp
L Device:C C4
U 1 1 618092CF
P 5225 1875
F 0 "C4" H 5340 1921 50  0000 L CNN
F 1 "10n" H 5340 1830 50  0000 L CNN
F 2 "Capacitor_THT:C_Disc_D3.0mm_W1.6mm_P2.50mm" H 5263 1725 50  0001 C CNN
F 3 "~" H 5225 1875 50  0001 C CNN
F 4 "C" H 5225 1875 50  0001 C CNN "Spice_Primitive"
F 5 "10n" H 5225 1875 50  0001 C CNN "Spice_Model"
F 6 "Y" H 5225 1875 50  0001 C CNN "Spice_Netlist_Enabled"
	1    5225 1875
	1    0    0    -1  
$EndComp
$Comp
L Device:C C5
U 1 1 61809B09
P 5225 2375
F 0 "C5" H 5340 2421 50  0000 L CNN
F 1 "10n" H 5340 2330 50  0000 L CNN
F 2 "Capacitor_THT:C_Disc_D3.0mm_W1.6mm_P2.50mm" H 5263 2225 50  0001 C CNN
F 3 "~" H 5225 2375 50  0001 C CNN
F 4 "C" H 5225 2375 50  0001 C CNN "Spice_Primitive"
F 5 "10n" H 5225 2375 50  0001 C CNN "Spice_Model"
F 6 "Y" H 5225 2375 50  0001 C CNN "Spice_Netlist_Enabled"
	1    5225 2375
	1    0    0    -1  
$EndComp
Wire Wire Line
	5225 2225 4575 2225
Wire Wire Line
	4575 2525 5225 2525
Wire Wire Line
	5225 2525 5575 2525
Connection ~ 5225 2525
Wire Wire Line
	4575 2725 5225 2725
Wire Wire Line
	5225 2725 5225 2925
$Comp
L Connector:Screw_Terminal_01x02 J1
U 1 1 617D4C3F
P 7475 3025
F 0 "J1" H 7555 3017 50  0000 L CNN
F 1 "Rails" H 7555 2926 50  0000 L CNN
F 2 "TerminalBlock_Phoenix:TerminalBlock_Phoenix_MKDS-1,5-2_1x02_P5.00mm_Horizontal" H 7475 3025 50  0001 C CNN
F 3 "~" H 7475 3025 50  0001 C CNN
	1    7475 3025
	1    0    0    -1  
$EndComp
Text GLabel 5575 1725 2    50   Input ~ 0
Rails0
Text GLabel 5575 2525 2    50   Input ~ 0
Rails1
Text GLabel 7275 3125 0    50   Input ~ 0
Rails0
Text GLabel 7275 3025 0    50   Input ~ 0
Rails1
Text GLabel 1725 2625 0    50   Input ~ 0
3V3_RPi
Text GLabel 1775 2175 0    50   Input ~ 0
GPIO27_RPi
Text GLabel 3125 2025 0    50   Input ~ 0
5V_RPi
Text GLabel 3175 2325 0    50   Input ~ 0
GPIO17_RPi
$Comp
L Connector:Conn_01x05_Male J3
U 1 1 617F4C97
P 7000 1875
F 0 "J3" H 7108 2256 50  0000 C CNN
F 1 "RPi Connect" H 7108 2165 50  0000 C CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x05_P2.54mm_Vertical" H 7000 1875 50  0001 C CNN
F 3 "~" H 7000 1875 50  0001 C CNN
	1    7000 1875
	1    0    0    -1  
$EndComp
Text GLabel 2425 1525 0    50   Input ~ 0
+18V
Wire Wire Line
	2425 1525 2575 1525
Text GLabel 7225 2750 0    50   Input ~ 0
+18V
$Comp
L power:GND #PWR0103
U 1 1 6180150B
P 7200 2075
F 0 "#PWR0103" H 7200 1825 50  0001 C CNN
F 1 "GND" H 7205 1902 50  0000 C CNN
F 2 "" H 7200 2075 50  0001 C CNN
F 3 "" H 7200 2075 50  0001 C CNN
	1    7200 2075
	1    0    0    -1  
$EndComp
Text GLabel 7200 1975 2    50   Input ~ 0
5V_RPi
Text GLabel 7200 1875 2    50   Input ~ 0
GPIO17_RPi
Connection ~ 2575 1525
Wire Wire Line
	2575 1525 3875 1525
Wire Wire Line
	2575 1525 2575 2925
Wire Wire Line
	4575 1725 5225 1725
Connection ~ 5225 1725
Wire Wire Line
	5225 1725 5575 1725
Wire Notes Line
	6325 850  6325 3875
Wire Notes Line
	900  3875 900  850 
Wire Notes Line
	8700 3875 8700 850 
Wire Notes Line
	900  850  8700 850 
Wire Notes Line
	900  3875 8700 3875
Text GLabel 7200 1775 2    50   Input ~ 0
3V3_RPi
Text GLabel 7200 1675 2    50   Input ~ 0
GPIO27_RPi
$Comp
L power:GND #PWR0102
U 1 1 617EB976
P 7100 2450
F 0 "#PWR0102" H 7100 2200 50  0001 C CNN
F 1 "GND" H 7105 2277 50  0000 C CNN
F 2 "" H 7100 2450 50  0001 C CNN
F 3 "" H 7100 2450 50  0001 C CNN
	1    7100 2450
	1    0    0    -1  
$EndComp
$Comp
L Connector:Screw_Terminal_01x02 J2
U 1 1 617E2F29
P 7425 2650
F 0 "J2" H 7505 2642 50  0000 L CNN
F 1 "Power" H 7505 2551 50  0000 L CNN
F 2 "TerminalBlock_Phoenix:TerminalBlock_Phoenix_MKDS-1,5-2_1x02_P5.00mm_Horizontal" H 7425 2650 50  0001 C CNN
F 3 "~" H 7425 2650 50  0001 C CNN
	1    7425 2650
	1    0    0    -1  
$EndComp
Wire Wire Line
	7225 2650 7225 2450
Wire Wire Line
	7225 2450 7100 2450
$Comp
L Mechanical:MountingHole H3
U 1 1 617D9FC3
P 6600 3600
F 0 "H3" H 6700 3646 50  0000 L CNN
F 1 "MountingHole" H 6700 3555 50  0000 L CNN
F 2 "MountingHole:MountingHole_2.7mm_M2.5" H 6600 3600 50  0001 C CNN
F 3 "~" H 6600 3600 50  0001 C CNN
	1    6600 3600
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H1
U 1 1 617DA2CF
P 6600 3400
F 0 "H1" H 6700 3446 50  0000 L CNN
F 1 "MountingHole" H 6700 3355 50  0000 L CNN
F 2 "MountingHole:MountingHole_2.7mm_M2.5" H 6600 3400 50  0001 C CNN
F 3 "~" H 6600 3400 50  0001 C CNN
	1    6600 3400
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H2
U 1 1 617DA6E2
P 7300 3400
F 0 "H2" H 7400 3446 50  0000 L CNN
F 1 "MountingHole" H 7400 3355 50  0000 L CNN
F 2 "MountingHole:MountingHole_2.7mm_M2.5" H 7300 3400 50  0001 C CNN
F 3 "~" H 7300 3400 50  0001 C CNN
	1    7300 3400
	1    0    0    -1  
$EndComp
$Comp
L Mechanical:MountingHole H4
U 1 1 617DAAF5
P 7300 3600
F 0 "H4" H 7400 3646 50  0000 L CNN
F 1 "MountingHole" H 7400 3555 50  0000 L CNN
F 2 "MountingHole:MountingHole_2.7mm_M2.5" H 7300 3600 50  0001 C CNN
F 3 "~" H 7300 3600 50  0001 C CNN
	1    7300 3600
	1    0    0    -1  
$EndComp
$EndSCHEMATC
