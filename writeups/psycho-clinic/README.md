# RuCTF2022 | Psycho clinic
## Description
You are an owner and at the same time a patient in clinic for the mentally ill. You can register new patients and clinic staff. Every patient can enter into contracts with doctors. Also they may ask for procedures. But procedure can be performed only when patient has active contract with doctor.

# Vuln
Flag is located in ```Doctor```'s description. Every time ```Doctor``` performs ```Procedure```, he writes it to report. Your aim is to prescribe to yourself a ```Procedure``` from needed ```Doctor```. It is expected that you can't create a ```Contract``` with an arbitrary ```Doctor``` because you do not know ```Doctor.Signature```. Therefore, you can't prescribe a procedure from this doctor. But ```Contract```'s comparasion [uses](../../services/psycho-clinic/src/psycho-clinic/Models/Contract.cs#L19) identity rendering. And there is a bug in interpolated string rendering:
```C#
var value = 42;
Renderer.Render($"{value}"); // -> "{42}"
Renderer.Render($"{value.ToString()}"); // -> "{value.ToString()}"
``` 

# Attack 
* Fetch needed ```doctorId``` from ```doctors/```
* Create stub ```Contract``` with some ```Doctor```. Save the ```contractId``` and ```expired```.
* Construct fake ```Contract``` using ```contractId, expired``` and needed ```doctorId```.
* Use fake ```Contract``` to prescribe some ```Procedure``` from ```Doctor``` with ```doctorId```.
* Perform this ```Procedure``` to get a flag :)

You can find sploit [here](../../sploits/psycho-clinic/sploit.py)

# Defense
Downgrade to .NET5.0 or remove ```.ToString()``` from rendering ```Contract```'s identity.
