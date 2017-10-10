package seq

/* python
class OrfFinder(object):

    ORF = namedtuple('ORF', 'start end')

    def __call__(self, seq):
        orfs = []

        for match in re.finditer('ATG', seq):
            # Find a start codon
            start = match.start()

            end = 0
            for i, codon in enumerate(codons(seq[start:])):
                if codon in STOP_CODONS:
                    # A stop codon was found,
                    # set the end position and return the ORF
                    end = start + i * 3 + 3
                    orfs.append(self.ORF(start + 1, end))
                    break

        return orfs

find_orfs = OrfFinder()
*/
