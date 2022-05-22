#include <ctype.h>

#include "diag.h"
#include "sha3.h"

uint64_t drugs_count = 69;
char * drugs[] = {
	"agomelatine",
	"alprazolam",
	"amisulpride",
	"amitriptyline",
	"aripiprazole",
	"asenapine",
	"benperidol",
	"buspirone",
	"carbamazepine",
	"cariprazine",
	"chlordiazepoxide",
	"chlorpromazine",
	"citalopram",
	"clomethiazole",
	"clomipramine",
	"clozapine",
	"diazepam",
	"diphenhydramine",
	"dosulepin",
	"doxepin",
	"duloxetine",
	"escitalopram",
	"fluoxetine",
	"flupentixol",
	"fluvoxamine",
	"haloperidol",
	"imipramine",
	"isocarboxazid",
	"lamotrigine",
	"levomepromazine",
	"lithium",
	"lofepramine",
	"loprazolam",
	"lorazepam",
	"lormetazepam",
	"lurasidone",
	"melatonin",
	"mianserin",
	"mirtazapine",
	"moclobemide",
	"nitrazepam",
	"nortriptyline",
	"olanzapine",
	"oxazepam",
	"paliperidone",
	"paroxetine",
	"pericyazine",
	"phenelzine",
	"pimozide",
	"pregabalin",
	"prochlorperazine",
	"promazine",
	"promethazine",
	"quetiapine",
	"reboxetine",
	"risperidone",
	"sertraline",
	"sulpiride",
	"temazepam",
	"tranylcypromine",
	"trazodone",
	"trifluoperazine",
	"trimipramine",
	"valproate",
	"venlafaxine",
	"vortioxetine",
	"zolpidem",
	"zopiclone",
	"zuclopenthixol"
};

void prescribe(char* diag, char* meds) {
	DEBUG("!! prescribe for '%s'\n", diag);

	uint64_t h = hash(diag, strlen(diag));
	strcpy(meds, drugs[h % drugs_count]);
}