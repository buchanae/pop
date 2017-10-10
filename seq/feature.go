package seq

/*
def get_transcript_seq(genome_sequence, exons, reverse=False):
    '''Exons are expected to be 1-based closed intervals'''

    exons = sorted(exons, key=lambda e: e.start)
    seq = ''.join(genome_sequence[exon.start - 1:exon.end] for exon in exons)
    if reverse:
        seq = reverse_complement(seq)
    return seq
*/
